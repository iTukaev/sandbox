package start

import (
	"context"
	"github.com/go-chi/chi"
	"log"
	"net"
	"net/http"
	"sandbox/Mod30/pkg/groupServise"
	"sandbox/Mod30/pkg/handlers/ageupdate"
	"sandbox/Mod30/pkg/handlers/create"
	"sandbox/Mod30/pkg/handlers/deleteuser"
	"sandbox/Mod30/pkg/handlers/getall"
	"sandbox/Mod30/pkg/handlers/getfriends"
	"sandbox/Mod30/pkg/handlers/makefriend"
	"time"
)

func Server(ctx context.Context, r *chi.Mux) {

	group := groupServise.NewService()
	r.Post("/create", create.NewHandler(group))
	r.Post("/make_friends", makefriend.NewHandler(group))
	r.Delete("/user", deleteuser.NewHandler(group))
	r.Get("/friends/{userId}", getfriends.NewHandle(group))
	r.Put("/{userId}", ageupdate.NewHandle(group))
	r.Get("/get", getall.NewHandle(group))


	listener, err := net.Listen("tcp", "127.0.0.1:8000")
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
		return
	}
	log.Println("firstServer stopped")
}
