package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sandbox/Mod31/pkg/db"
	"sandbox/Mod31/pkg/httprouter"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	HandleSignals(cancel)
	db.MongoDbStart()

	httprouter.StartHttpServer(ctx, db.Client)

	log.Println("All services stopped. Goodbye!")
}

func HandleSignals(cancel context.CancelFunc) {
	go func() {
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
		<- exitCh
		cancel()
	}()
}