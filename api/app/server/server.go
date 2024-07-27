package server

import (
	gen "api/openapi"

	"github.com/labstack/echo/v4"
)

var _ gen.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetTest(ctx echo.Context) error {
	ctx.Logger().Info("GetTest")
	return nil
}
