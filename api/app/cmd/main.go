package main

import (
	"api/app/config"
	"api/app/server"
	"fmt"
	"log"

	gen "api/openapi"

	mw "api/app/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer()

	e := echo.New()
	e.Use(mw.Logger(mw.SlogLoggerFunc))

	gen.RegisterHandlers(e, server)

	log.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", c.Port)))
}
