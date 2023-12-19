package app

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/config"
	"github.com/smartfood-capstone/backend/internal/database"
	"github.com/smartfood-capstone/backend/internal/foods"
	"github.com/smartfood-capstone/backend/internal/history"
	"github.com/smartfood-capstone/backend/internal/predict"
	"github.com/smartfood-capstone/backend/internal/routes"
	"github.com/smartfood-capstone/backend/internal/server"
	"github.com/smartfood-capstone/backend/internal/shops"
)

type StartCmd struct {
	Logger *logrus.Logger
	Server server.Server
}

func NewStartCmd() *StartCmd {
	l := logrus.New()
	cfg := config.New()
	db := database.New(cfg)

	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			fileName := fmt.Sprintf("%s:%d", f.File, f.Line)
			arrOfFunc := strings.Split(f.Function, "/")
			funcName := arrOfFunc[len(arrOfFunc)-1]

			return funcName, fileName
		},
	})

	svr := server.New(l, cfg)
	routes.InitRoutes(svr)

	foodsRepository := foods.NewRepository(db, l)
	foodsService := foods.NewService(foodsRepository, l)
	foodsController := foods.NewController(foodsService, l)
	foods.RegisterRoute(svr, foodsController)

	historyRepository := history.NewRepository(db, l)
	historyService := history.NewService(historyRepository, l)
	historyController := history.NewController(l, historyService)
	history.RegisterRoute(svr, historyController)

	predictRepository := predict.NewRepository(db, l)
	predictService := predict.NewService(predictRepository, l)
	predictController := predict.NewController(predictService, l)
	predict.RegisterRoute(svr, predictController)

	shopsRepository := shops.NewRepository(db, l)
	shopsService := shops.NewService(shopsRepository, l)
	shopsController := shops.NewController(shopsService, l)
	shops.RegisterRoute(svr, shopsController)

	return &StartCmd{
		Server: svr,
		Logger: l,
	}
}

func (s *StartCmd) Start() {
	if err := s.Server.Start(); err != nil {
		logrus.Fatal("failed to start server")
	}
}
