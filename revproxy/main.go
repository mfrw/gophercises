package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewProxy(target string) *Proxy {
	url, _ := url.Parse(target)
	return &Proxy{
		target: url,
		proxy:  httputil.NewSingleHostReverseProxy(url),
	}
}

func (p *Proxy) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	p.proxy.ServeHTTP(w, r)
}

var (
	port        *string
	redirecturl *string
)

func main() {
	const (
		defaultPort        = ":9090"
		defaultPortUsage   = "default server port, :9090"
		defaultTarget      = "socks5://localhost:9050"
		defaultTargetUsage = "default redirect url (Tor), socks5://localhost:9050"
	)
	port = flag.String("port", defaultPort, defaultPortUsage)
	redirecturl = flag.String("url", defaultTarget, defaultTargetUsage)

	flag.Parse()

	fmt.Println("server will run on :", *port)
	fmt.Println("redirecting to :", *redirecturl)

	// proxy
	proxy := NewProxy(*redirecturl)

	http.HandleFunc("/proxyServer", ProxyServer)

	// server redirection
	http.HandleFunc("/", proxy.handle)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

}

func ProxyServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Reverse proxy server running. Accepting at port: ", +*port+" Redirecting to: "+*redirecturl))
}
