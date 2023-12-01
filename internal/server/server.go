package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server interface {
	Start(port string)
	App() *echo.Echo
}

type server struct {
	app *echo.Echo
}

func New() Server {
	app := echo.New()
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())

	return &server{
		app: app,
	}
}

func (s *server) Start(port string) {
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	s.app.Start(addr)
}

func (s *server) App() *echo.Echo {
	return s.app
}
