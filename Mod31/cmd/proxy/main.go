package main

import (
	"context"
	"log"
	"sandbox/Mod31/pkg/proxy/start"
	"sandbox/Mod31/pkg/proxy/stop"
)

const (
	proxyHost = "127.0.0.1:8000"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	stop.Proxy(cancel)
	start.Proxy(ctx, proxyHost)
	log.Println("main", proxyHost, "done")
}

