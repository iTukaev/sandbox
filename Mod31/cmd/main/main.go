package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"sandbox/Mod30/pkg/server/start"
	"sandbox/Mod30/pkg/server/stop"
)

func main() {
	router := chi.NewRouter()
	log.Println("router created")
	router.Use(middleware.Logger)
	log.Println("logger started")

	ctx, cancel := context.WithCancel(context.Background())
	stop.Server(cancel)
	start.Server(ctx, router)

	log.Println("main: done")
}