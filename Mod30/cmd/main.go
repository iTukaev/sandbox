package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"sandbox/Mod30/pkg/server/start"
	"sandbox/Mod30/pkg/server/stop"
)

const(
	ADDRESS = "127.0.0.1:8000"
)

func main() {
	router := chi.NewRouter()
	log.Println("router created")
	router.Use(middleware.Logger)
	log.Println("logger started")

	ctx, cancel := context.WithCancel(context.Background())
	stop.Server(cancel)
	start.Server(ctx, router, ADDRESS)

	log.Println("main: done")
}