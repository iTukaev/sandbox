package internal

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"sandbox/Mod30/pkg/entity"
	"sandbox/Mod30/pkg/service"
)

func StartHttpServer(ctx context.Context) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	dpt := entity.NewDepartment()
	createUser := service.NewCreateUser(dpt)

	srv := &http.Server{Addr: "localhost:8000", Handler: r}
	log.Println("server started")

	r.Post("/create", createUser.Create)
	r.Post("/make_friends", createUser.MakeFriend)
	r.Delete("/user", createUser.DeleteUser)
	r.Get("/friends/{userId}", createUser.Friends)
	r.Put("/{userId}", createUser.AgeUpdate)
	r.Get("/get", createUser.GetAllUsers)

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