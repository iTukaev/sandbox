package firstStart

import (
	"context"
	"github.com/go-chi/chi"
	"log"
	"net"
	"net/http"
	"sandbox/Mod31/pkg/db/dbService"
	"sandbox/Mod31/pkg/handlers/ageupdate"
	"sandbox/Mod31/pkg/handlers/create"
	"sandbox/Mod31/pkg/handlers/deleteuser"
	"sandbox/Mod31/pkg/handlers/getall"
	"sandbox/Mod31/pkg/handlers/getfriends"
	"sandbox/Mod31/pkg/handlers/makefriend"
	"time"
)

func Server(ctx context.Context, r *chi.Mux) {

	client := dbService.NewService()
	r.Post("/create", create.NewHandler(client))
	r.Post("/make_friends", makefriend.NewHandler(client))
	r.Delete("/user", deleteuser.NewHandler(client))
	r.Get("/friends/{userId}", getfriends.NewHandle(client))
	r.Put("/{userId}", ageupdate.NewHandle(client))
	r.Get("/get", getall.NewHandle(client))


	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Println(err)
	}

	server := &http.Server{
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("firstServer started")
	go func() {
		log.Println(server.Serve(listener))
	}()

	<-ctx.Done()
	if err := server.Close(); err != nil {
		log.Println(err)
	}
	if err := dbService.Stop(ctx); err != nil {
		log.Println(err)
	}

	log.Println("firstServer stopped")
}
