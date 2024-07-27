package main

import (
	"api/app/config"
	"api/app/server"
	"fmt"
	"log"

	gen "api/openapi"

	"github.com/labstack/echo/v4"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer()

	e := echo.New()

	gen.RegisterHandlers(e, server)

	log.Fatal(e.Start(fmt.Sprintf("localhost:%s", c.Port)))
}
