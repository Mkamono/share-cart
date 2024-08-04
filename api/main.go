package main

import (
	"api/app/config"
	h "api/app/handler"
	"api/app/infra/jwt"
	db "api/app/infra/repository"
	mw "api/app/middleware"
	"api/app/oas"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	const auth0SubjectKey = "subject"

	// Middleware
	loggerMiddleware := mw.NewSlogLogger()
	jwtMiddleware, err := mw.NewJwtMiddleware(c.AUTH0.Domain, c.AUTH0.Audience, auth0SubjectKey)
	if err != nil {
		log.Fatal(err)
	}

	// Handler
	jwtClient := jwt.NewJwtClient(auth0SubjectKey)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	db, err := db.New(c.GetDBConnStr())
	if err != nil {
		log.Fatal(err)
	}

	service := h.NewHandler(
		jwtClient,
		logger,
		db,
	)

	// Server
	srv, err := oas.NewServer(
		service,
		oas.WithMiddleware(loggerMiddleware, jwtMiddleware),
	)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%s", c.Port)
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatal(err)
	}
}
