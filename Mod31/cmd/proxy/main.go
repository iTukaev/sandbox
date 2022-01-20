package main

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var count int

type Config struct {
	Proxy string `yaml:"proxy"`
	FirstServer string `yaml:"first_server"`
	SecondServer string `yaml:"second_server"`
}

func main() {
	cfg := new(Config)
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		log.Fatalf("Config not available %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	stop(cancel)
	start(ctx, *cfg)
	log.Printf("main %s done", cfg.Proxy)
}


func start(ctx context.Context, cfg Config) {
	listener, err := net.Listen("tcp", cfg.Proxy)
	if err != nil {
		log.Fatalf("Listener start error: %v", err)
	}

	server := &http.Server{
		Handler: nil,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		log.Println(server.Serve(listener))
	}()

	log.Printf("proxy %s started", cfg.Proxy)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ReversProxyStarter(w, r, cfg)
	})

	<-ctx.Done()
	if err := server.Close(); err != nil {
		log.Printf("Server closing error: %v", err)
	}
	log.Printf("proxy %s stopped", cfg.Proxy)
}


func stop(cancel context.CancelFunc) {
	go func() {
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
		<- exitCh
		cancel()
	}()
}

func ReversProxyStarter(w http.ResponseWriter, r *http.Request, cfg Config) {
	firstURL, err := url.Parse("http://" + cfg.FirstServer)
	if err != nil {
		log.Printf("URL %s parsing error: %v", cfg.FirstServer, err)
		return
	}

	secondURL, err := url.Parse("http://" + cfg.SecondServer)
	if err != nil {
		log.Printf("URL %s parsing error: %v", cfg.SecondServer, err)
		return
	}

	if count == 0 {
		log.Printf("redirected on: %s", cfg.FirstServer)
		proxy := httputil.NewSingleHostReverseProxy(firstURL)
		proxy.ServeHTTP(w, r)
		count++
		return
	}

	if count == 1 {
		log.Printf("redirected on: %s", cfg.SecondServer)
		proxy := httputil.NewSingleHostReverseProxy(secondURL)
		proxy.ServeHTTP(w, r)
		count--
		return
	}
}