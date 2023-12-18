package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/config"
	"github.com/smartfood-capstone/backend/internal/routes"
)

type Server interface {
	Start() error
	App() *echo.Echo
	Log() *logrus.Logger
}

type server struct {
	app    *echo.Echo
	logger *logrus.Logger
	config config.Config
}

func New(l *logrus.Logger, cfg config.Config) Server {
	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Recover())
	app.Logger.SetOutput(l.Writer())
	app.Use(middleware.Logger())

	routes.InitRoutes(l, app)

	return &server{
		app:    app,
		logger: l,
		config: cfg,
	}
}

func (s *server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	return s.app.Start(addr)
}

func (s *server) App() *echo.Echo {
	return s.app
}

func (s *server) Log() *logrus.Logger {
	return s.logger
}
