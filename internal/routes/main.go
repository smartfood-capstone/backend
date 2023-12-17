package routes

// Define Echo routes

import (
	"github.com/labstack/echo/v4"
	"github.com/smartfood-capstone/backend/internal/controller"
)

func InitRoutes(e *echo.Echo) {
	c := controller.New()

	v1 := e.Group("/api/v1")

	v1.GET("/health", c.HealthCheck)
}
