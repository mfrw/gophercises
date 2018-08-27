package ex3

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:] // Trime the leading /

	if chapter, ok := h.s[path]; ok {
		tmpl := template.Must(template.New("").Parse(htmlTemplate))
		err := tmpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong ..", http.StatusNotFound)
		}
	}
}

func JsonStory(r io.Reader) (Story, error) {
	var story Story

	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		if err != nil {
			return nil, err
		}
	}
	return story, nil
}
