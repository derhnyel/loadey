package main

// TODO: spin up server instance
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	proxy       *httputil.ReverseProxy
	URLString   string
	isAvailable bool
}

type loadBalancer struct {
	servers         []*server
	roundRobinCount int
}

func (l *loadBalancer) getAvailableServer() *server {
	server := l.servers[l.roundRobinCount%len(l.servers)]
	for !server.isAvailable {
		l.roundRobinCount++
		server = l.servers[l.roundRobinCount%len(l.servers)]
	}
	l.roundRobinCount++
	return server
}

func (l *loadBalancer) directToServer(rw http.ResponseWriter, r *http.Request) {
	srv := l.getAvailableServer()
	fmt.Println("Directing to server ", srv)
	srv.proxy.ServeHTTP(rw, r)
	fmt.Println("Served ", srv)
}

func (l *loadBalancer) addServers(srvs ...*server) {
	l.servers = append(l.servers, srvs...)
}

func NewServer(urlString string) (*server, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	return &server{proxy, urlString, true}, nil
}

func main() {
	s1, _ := NewServer("https://www.facebook.com")
	s2, _ := NewServer("http://www.bing.com")
	s3, _ := NewServer("http://www.duckduckgo.com")
	s4, _ := NewServer("https://www.github.com")

	l := &loadBalancer{roundRobinCount: 0}
	l.addServers(s1, s2, s3, s4)
	fmt.Println("The loadbalancer is", l)
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("receiving request from /")
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Add("Pragma", "no-cache")
		w.Header().Add("Expires", "0")
		l.directToServer(w, r)
	}
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8000", nil)

}
