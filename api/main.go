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

	const auth0SubjectKey = "subject"

	service := h.NewHandler(auth0SubjectKey)
	logger := mw.NewSlogLogger()
	jwtMiddleware, err := mw.NewJwtMiddleware(c.AUTH0.Domain, c.AUTH0.Audience, auth0SubjectKey)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := oas.NewServer(service, oas.WithMiddleware(logger), oas.WithMiddleware(jwtMiddleware))
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%s", c.Port)
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatal(err)
	}
}
