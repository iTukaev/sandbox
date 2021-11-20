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
	secondInstanceHost = "127.0.0.1:8002"
)

func main() {
	router := chi.NewRouter()
	log.Println("router " + secondInstanceHost + " created")
	router.Use(middleware.Logger)
	log.Println("logger " + secondInstanceHost + " started")

	ctx, cancel := context.WithCancel(context.Background())
	stop.Server(cancel)
	start.Server(ctx, router, secondInstanceHost)

	log.Println("main " + secondInstanceHost + " done")
}