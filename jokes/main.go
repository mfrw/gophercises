package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/publicsuffix"
)

type joke struct {
	Type  string `json:"type"`
	Value struct {
		Categories []string `json:"categories"`
		ID         int64    `json:"id"`
		Joke       string   `json:"joke"`
	} `json:"value"`
}

type JokeClient struct {
	client  *http.Client
	baseUrl string
}

func NewJokeClient() *JokeClient {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		jar = nil
	}
	return &JokeClient{
		client: &http.Client{
			Jar:     jar,
			Timeout: 15 * time.Second,
		},
		baseUrl: "http://api.icndb.com/jokes/random",
	}
}

func (j *JokeClient) RandomJson() ([]byte, error) {
	req, err := http.NewRequest("GET", j.baseUrl, nil)
	if err != nil {
		return []byte(""), err
	}

	resp, err := j.client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	msg, err := ioutil.ReadAll(resp.Body)
	return msg, nil
}

func (j *JokeClient) Random() (string, error) {
	req, err := http.NewRequest("GET", j.baseUrl, nil)
	if err != nil {
		return "", err
	}

	resp, err := j.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	nj := new(joke)
	err = json.NewDecoder(resp.Body).Decode(nj)
	if err != nil {
		return "", err
	}
	return nj.Value.Joke, nil
}

var promInst = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "joke_handler_in_flight",
	Help: "Inflight handlers serving clients",
})
var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "joke_hadler_service_time",
	Help:    "Time taken to service",
	Buckets: []float64{0.1, 0.2, 0.21, 0.22, 0.23, 0.24, 0.25, 0.26, 0.27, 0.28, 0.29, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0},
}, []string{"code", "method"})

func init() {
	prometheus.MustRegister(promInst)
	prometheus.MustRegister(requestDuration)
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	promInst.Inc()
	defer promInst.Dec()

	jc := NewJokeClient()
	joke, err := jc.RandomJson()
	if err != nil {
		http.Error(w, "Interal Service Error", http.StatusInternalServerError)
		return
	}
	br := bytes.NewReader(joke)
	io.Copy(w, br)
}

func main() {
	jokeHandlerInst := promhttp.InstrumentHandlerDuration(requestDuration, http.HandlerFunc(jokeHandler))

	http.HandleFunc("/random", jokeHandlerInst)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}

func JokeRandom(n int) {
	jc := NewJokeClient()
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			j, err := jc.Random()
			if err != nil {
				return
			}
			fmt.Printf("%3d : %s\n", i+1, html.UnescapeString(j))
		}(i)
	}
	wg.Wait()
}
