package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sandbox/Mod30/pkg/groupServise"
	"sandbox/Mod30/pkg/handlers/create"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	HandleSignals(cancel)
	StartHttpServer(ctx)

	log.Println("main: done")
}

func StartHttpServer(ctx context.Context) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//group := groupServise.NewService()
	var h create.Handle
	h.GroupService = groupServise.NewService()

	srv := &http.Server{Addr: "localhost:8000", Handler: r}
	log.Println("server started")

	r.Post("/create", h.Create)
	r.Post("/make_friends", h.)
	//r.Delete("/user", createUser.DeleteUser)
	//r.Get("/friends/{userId}", createUser.Friends)
	//r.Put("/{userId}", createUser.AgeUpdate)
	//r.Get("/get", createUser.GetAllUsers)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("server will be stopped")
			CloseHttpServer(srv)
			return
		default:
			continue
		}
	}
}

func CloseHttpServer(srv *http.Server) {
	if err := srv.Close(); err != nil {
		log.Println(err)
	}
}

func HandleSignals(cancel context.CancelFunc) {
	go func() {
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
		<- exitCh
		cancel()
	}()
}