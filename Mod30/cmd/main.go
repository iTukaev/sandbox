package main

import (
	"context"
	"log"
	"sandbox/Mod30/cmd/internal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	internal.HandleSignals(cancel)
	internal.StartHttpServer(ctx)

	log.Println("main: done")
}

//func StartHttpServer(ctx context.Context) {
//	r := chi.NewRouter()
//	r.Use(middleware.Logger)
//
//	dpt := entity.NewDepartment()
//	createUser := service.NewCreateUser(dpt)
//
//	srv := &http.Server{Addr: "localhost:8000", Handler: r}
//	log.Println("server started")
//
//	r.Post("/create", createUser.Create)
//	r.Post("/make_friends", createUser.MakeFriend)
//	r.Delete("/user", createUser.DeleteUser)
//	r.Get("/friends/{userId}", createUser.Friends)
//	r.Put("/{userId}", createUser.AgeUpdate)
//	r.Get("/get", createUser.GetAllUsers)
//
//	go func() {
//		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
//			log.Fatalf("ListenAndServe(): %s", err)
//		}
//	}()
//
//	for {
//		select {
//		case <-ctx.Done():
//			log.Println("server will be stopped")
//			CloseHttpServer(srv)
//			return
//		default:
//			continue
//		}
//	}
//}

//func CloseHttpServer(srv *http.Server) {
//	if err := srv.Close(); err != nil {
//		log.Println(err)
//	}
//}

//func HandleSignals(cancel context.CancelFunc) {
//	go func() {
//		exitCh := make(chan os.Signal, 1)
//		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
//		<- exitCh
//		cancel()
//	}()
//}