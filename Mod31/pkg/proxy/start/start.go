package start

import (
	"context"
	"log"
	"net/http"
	"sandbox/Mod31/pkg/proxy/handler"
)

func Proxy(ctx context.Context, proxyHost string)  {
	go func() {
		if err := http.ListenAndServe(proxyHost, nil); err != nil {
			log.Println("proxy", proxyHost, "not started:", err)
			panic(err)
		}
	}()
	log.Println("proxy", proxyHost, "started")

	http.HandleFunc("/", handler.Proxy)

	<-ctx.Done()
	log.Println("proxy", proxyHost, "stopped")
}
