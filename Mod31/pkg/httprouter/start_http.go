package httprouter

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"sandbox/Mod31/pkg/db"
	"sandbox/Mod31/pkg/service"
)

func StartHttpServer(ctx context.Context) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	cl, err := db.MongoDbStart()
	if err != nil {
		log.Fatalln(err)
		return
	}

	dbHandler := service.NewMongoClient(cl)

	srv := &http.Server{Addr: "localhost:8000", Handler: r}
	log.Println("server started")

	r.Post("/create", dbHandler.Create)
	//r.Post("/make_friends", dbHandler.MakeFriend)
	r.Delete("/user", dbHandler.DeleteUser)
	//r.Get("/friends/{userId}", dbHandler.Friends)
	r.Put("/{userId}", dbHandler.AgeUpdate)
	r.Get("/get", dbHandler.GetAllUsers)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			CloseHttpServer(srv)
			db.MongoDbStop(cl.Client)
			return
		default:
			//здесь не знаю как корректно завершить все процессы сервера
			//всё, что я нашёл, быдло связано просто с временной задержкой в 3-5 секунд
			continue
		}
	}
}
