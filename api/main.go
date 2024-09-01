package main

import (
	"api/app/config"
	h "api/app/handler"
	db "api/app/infra/repository"
	mw "api/app/middleware"
	"api/app/oas"
	"api/app/shared/ctxlogger"
	jwtShared "api/app/shared/jwt"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"os"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", c.AUTH0.Domain))
	if err != nil {
		log.Fatal(err)
	}
	jwtClient, err := jwtShared.NewClient(issuerURL, []string{c.AUTH0.Audience})
	if err != nil {
		log.Fatal(err)
	}

	// logger
	baseHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: ctxlogger.GoogleCloudReplacer,
	})

	slogHandler := ctxlogger.NewHandler(baseHandler)
	slog.SetDefault(slog.New(slogHandler))

	// Middleware
	loggerMiddleware := mw.NewRequestLogger(c.ProjectID)

	// Handler
	db, err := db.New(c.GetDBConnStr())
	if err != nil {
		log.Fatal(err)
	}

	handler := h.NewHandler(jwtClient, db)
	securityHandler := h.NewSecurityHandler(jwtClient)

	// Server
	srv, err := oas.NewServer(
		handler,
		securityHandler,
		oas.WithMiddleware(loggerMiddleware),
	)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%s", c.Port)
	if err := http.ListenAndServe(addr, srv); err != nil {
		log.Fatal(err)
	}
}
