package main

import (
	"context"
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sandbox/Mod31/pkg/db/dbService"
	"sandbox/Mod31/pkg/handlers/ageupdate"
	"sandbox/Mod31/pkg/handlers/create"
	"sandbox/Mod31/pkg/handlers/deleteuser"
	"sandbox/Mod31/pkg/handlers/getall"
	"sandbox/Mod31/pkg/handlers/getfriends"
	"sandbox/Mod31/pkg/handlers/makefriend"
	"syscall"
	"time"
)

type Config struct {
	FirstServer string `yaml:"first_server"`
	SecondServer string `yaml:"second_server"`
}

func main() {
	cfg := new(Config)
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		panic(err)
	}

	var host = cfg.FirstServer
	second := flag.Bool("s", false, "if -s enable, second server start")
	flag.Parse()
	if *second {
		host = cfg.SecondServer
	}

	router := chi.NewRouter()
	log.Printf("router %s created", host)
	router.Use(middleware.Logger)
	log.Printf("logger %s created", host)

	ctx, cancel := context.WithCancel(context.Background())
	stop(cancel)
	start(ctx, router, host)

	log.Printf("main %s done", host)
}


func start(ctx context.Context, r *chi.Mux, address string) {

	client := dbService.NewService()
	r.Post("/create", create.NewHandler(client))
	r.Post("/make_friends", makefriend.NewHandler(client))
	r.Delete("/user", deleteuser.NewHandler(client))
	r.Get("/friends/{userId}", getfriends.NewHandler(client))
	r.Put("/{userId}", ageupdate.NewHandler(client))
	r.Get("/get", getall.NewHandler(client))


	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
	}

	server := &http.Server{
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Printf("server %s started", address)
	go func() {
		log.Println(server.Serve(listener))
	}()

	<-ctx.Done()
	if err := server.Close(); err != nil {
		log.Printf("Server closing error: %v", err)
	}
	if err := dbService.Stop(ctx); err != nil {
		log.Printf("MongoDB stop error: %v", err)
	}

	log.Printf("server %s stopped", address)
}


func stop(cancel context.CancelFunc) {
	go func() {
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
		<- exitCh
		cancel()
	}()
}
