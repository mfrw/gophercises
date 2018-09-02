package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mfrw/gophercises/ex11/hn"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	sc := storyCache{
		numStories: numStories,
		duration:   3 * time.Second,
	}
	go func() {
		tc := time.NewTicker(3 * time.Second)
		for {
			temp := storyCache{
				numStories: numStories,
				duration:   6 * time.Second,
			}
			temp.stories()
			sc.mutex.Lock()
			sc.cache = temp.cache
			sc.expiration = temp.expiration
			sc.mutex.Unlock()

			<-tc.C
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.stories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

type storyCache struct {
	numStories int
	cache      []item
	duration   time.Duration
	expiration time.Time
	mutex      sync.Mutex
}

func (sc *storyCache) stories() ([]item, error) {
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	if time.Now().Sub(sc.expiration) < 0 {
		return sc.cache, nil
	}
	stories, err := getTopStories(sc.numStories)
	if err != nil {
		return nil, err
	}
	cache = stories
	sc.expiration = time.Now().Add(sc.duration * time.Second)

	sc.cache = stories
	return sc.cache, nil
}

var (
	cache           []item
	cmtx            sync.Mutex
	cacheExpiration time.Time
)

func getCachedStories(numStories int) ([]item, error) {
	cmtx.Lock()
	defer cmtx.Unlock()
	if time.Now().Sub(cacheExpiration) < 0 {
		return cache, nil
	}

	stories, err := getTopStories(numStories)
	if err != nil {
		return nil, err
	}
	cache = stories
	cacheExpiration = time.Now().Add(5 * time.Second)

	return cache, nil
}

func getTopStories(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("Failed to load top stories")
	}
	var stories []item

	var wg sync.WaitGroup
	type result struct {
		item item
		err  error
		idx  int
	}
	restulCh := make(chan result)
	for idx, id := range ids {
		wg.Add(1)
		go func(idx, id int) {
			defer wg.Done()
			hnItem, err := client.GetItem(id)
			if err != nil {
				restulCh <- result{err: err}
			}
			restulCh <- result{item: parseHNItem(hnItem), idx: idx}
		}(idx, id)
	}
	go func() {
		wg.Wait()
		close(restulCh)
	}()

	var results []result
	for res := range restulCh {
		if res.err != nil {
			continue
		}

		if isStoryLink(res.item) {
			results = append(results, res)
			if len(results) >= numStories {
				break
			}
		}

	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx < results[j].idx
	})
	for i := 0; i < numStories; i++ {
		stories = append(stories, results[i].item)
	}
	return stories, nil
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
