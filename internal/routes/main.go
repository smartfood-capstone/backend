package routes

// Define Echo routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/controller"
)

func InitRoutes(l *logrus.Logger, e *echo.Echo) {
	c := controller.New(l)

	v1 := e.Group("/api/v1")

	v1.GET("/health", c.HealthCheck)

	// Food
	food := v1.Group("/food")
	food.GET("", c.GetAllFoods)
	food.GET("/:id", c.GetFoodDetail)

	// Shop
	shop := v1.Group("/shop")
	shop.GET("", c.GetAllShops)
	shop.GET("/:id", c.GetShopDetail)

	// User
	user := v1.Group("/user")
	user.GET("/:id/history", c.GetUserHistory)

	// Detect
	detect := v1.Group("/detect")
	detect.POST("", c.DetectFood)
}
