package main

import (
	"api/app/config"
	h "api/app/handler"
	mw "api/app/middleware"
	"api/app/oas"
	"fmt"
	"log"
	"net/http"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	service := h.NewHandler()
	logger := mw.NewSlogLogger()

	srv, err := oas.NewServer(service, oas.WithMiddleware(logger))
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%s", c.Port)
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatal(err)
	}
}
