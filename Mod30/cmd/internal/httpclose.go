package internal

import (
	"log"
	"net/http"
)

func CloseHttpServer(srv *http.Server) {
	if err := srv.Close(); err != nil {
		log.Println(err)
	}
}
