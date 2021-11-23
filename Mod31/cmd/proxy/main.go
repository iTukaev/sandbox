package main

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

var count int

type Config struct {
	Proxy string `yaml:"proxy"`
	FirstServer string `yaml:"first_server"`
	SecondServer string `yaml:"second_server"`
}

func main() {
	cfg := new(Config)
	if err := cleanenv.ReadConfig("./Mod31/config/config.yml", cfg); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	stop(cancel)
	start(ctx, *cfg)
	log.Println("main", cfg.Proxy, "done")
}


func start(ctx context.Context, cfg Config)  {
	go func() {
		if err := http.ListenAndServe(cfg.Proxy, nil); err != nil {
			log.Println("proxy", cfg.Proxy, "not started:", err)
			panic(err)
		}
	}()
	log.Println("proxy", cfg.Proxy, "started")

	http.HandleFunc("/", cfg.Handler)

	<-ctx.Done()
	log.Println("proxy", cfg.Proxy, "stopped")
}


func stop(cancel context.CancelFunc) {
	go func() {
		exitCh := make(chan os.Signal, 1)
		signal.Notify(exitCh, os.Interrupt, syscall.SIGINT)
		<- exitCh
		cancel()
	}()
}

func (cfg Config) Handler(w http.ResponseWriter, r *http.Request, ) {
	firstURL, err := url.Parse("http://" + cfg.FirstServer)
	if err != nil {
		log.Println(cfg.FirstServer, "URL error: ", err)
		return
	}

	secondURL, err := url.Parse("http://" + cfg.SecondServer)
	if err != nil {
		log.Println(cfg.SecondServer, "URL error: ", err)
		return
	}

	if count == 0 {
		log.Println("redirected on:", cfg.FirstServer)
		proxy := httputil.NewSingleHostReverseProxy(firstURL)
		proxy.ServeHTTP(w, r)
		count++
		return
	}

	if count == 1 {
		log.Println("redirected on:", cfg.SecondServer)
		proxy := httputil.NewSingleHostReverseProxy(secondURL)
		proxy.ServeHTTP(w, r)
		count--
		return
	}
}