package foods

import (
	"github.com/smartfood-capstone/backend/internal/server"
)

func RegisterRoute(s server.Server, c controller) {
	api := s.App().Group("/api")
	v1 := api.Group("/v1")

	foods := v1.Group("/foods")
	foods.GET("", c.GetAllFoods)
	foods.GET("/:id", c.GetFoodDetail)
}
