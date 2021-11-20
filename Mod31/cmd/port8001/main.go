package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"sandbox/Mod31/pkg/server/start"
	"sandbox/Mod31/pkg/server/stop"
)

const (
	firstInstanceHost = "127.0.0.1:8001"
)

func main() {
	router := chi.NewRouter()
	log.Println("router " + firstInstanceHost + " created")
	router.Use(middleware.Logger)
	log.Println("logger " + firstInstanceHost + " started")

	ctx, cancel := context.WithCancel(context.Background())
	stop.Server(cancel)
	start.Server(ctx, router, firstInstanceHost)

	log.Println("main " + firstInstanceHost + " done")
}