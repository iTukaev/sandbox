package handler

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	firstInstanceHost = "http://127.0.0.1:8001/"
	secondInstanceHost = "http://127.0.0.1:8002/"
)

var count int

func Proxy(w http.ResponseWriter, r *http.Request, ) {
	firstURL, err := url.Parse(firstInstanceHost)
	if err != nil {
		log.Println(firstInstanceHost, "URL error: ", err)
		return
	}

	secondURL, err := url.Parse(secondInstanceHost)
	if err != nil {
		log.Println(secondInstanceHost, "URL error: ", err)
		return
	}

	if count == 0 {
		log.Println("redirected on:", firstInstanceHost)
		proxy := httputil.NewSingleHostReverseProxy(firstURL)
		proxy.ServeHTTP(w, r)
		count++
		return
	}

	if count == 1 {
		log.Println("redirected on: ", secondInstanceHost)
		proxy := httputil.NewSingleHostReverseProxy(secondURL)
		proxy.ServeHTTP(w, r)
		count--
		return
	}
}
