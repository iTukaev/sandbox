package proxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	proxyHost = "http://127.0.0.1:8000/"
	firstInstanceHost = "http://127.0.0.1:8001/"
	secondInstanceHost = "http://127.0.0.1:8002/"
)

var count int

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Host)

	firstURL, err := url.Parse(firstInstanceHost)
	if err != nil {
		return
	}

	secondURL, err := url.Parse(secondInstanceHost)
	if err != nil {
		return
	}

	if count == 0 {
		log.Println(r.URL.Host)
		proxy := httputil.NewSingleHostReverseProxy(firstURL)
		proxy.ServeHTTP(w, r)
		count++
		return
	}

	if count == 1 {
		log.Println(r.URL.Host)
		proxy := httputil.NewSingleHostReverseProxy(secondURL)
		proxy.ServeHTTP(w, r)
		count--
		return
	}
}

func Start(ctx context.Context) {
	log.Println("Proxy started")
	http.HandleFunc("/", handlerProxy)
	go func() {
		if err := http.ListenAndServe(proxyHost, nil); err != nil {
			log.Println(err)
			panic(err)
		}
	}()

	<-ctx.Done()
}