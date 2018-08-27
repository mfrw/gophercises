package ex3

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	tmpl = template.Must(template.New("").Parse(htmlTemplate))
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
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{
		s:      s,
		t:      tmpl,
		pathFn: defaultPathFn,
	}
	for _, o := range opts {
		o(&h)
	}
	return h
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:] // Trime the leading /
}

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)
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
