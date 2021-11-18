package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"sandbox/Mod31/pkg/firstServer/firstStart"
	"sandbox/Mod31/pkg/firstServer/firstStop"
)

func main() {
	router := chi.NewRouter()
	log.Println("router created")
	router.Use(middleware.Logger)
	log.Println("logger started")

	ctx, cancel := context.WithCancel(context.Background())
	firstStop.Server(cancel)
	firstStart.Server(ctx, router)
	//secondStart.Server(ctx, router)
	//proxy.Start(ctx)

	log.Println("main: done")
}