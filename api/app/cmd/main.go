package main

import (
	"api/app/config"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	log.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", c.Port)))
}
