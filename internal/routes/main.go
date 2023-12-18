package routes

// Define Echo routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartfood-capstone/backend/internal/server"
	"github.com/smartfood-capstone/backend/internal/util"
)

func HealthCheck(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func InitRoutes(s server.Server, e *echo.Echo) {
	v1 := e.Group("/api/v1")

	v1.GET("/health", HealthCheck)
}
