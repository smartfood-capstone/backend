package foods

import (
	"github.com/smartfood-capstone/backend/internal/server"
)

func RegisterRoute(s server.Server, c IController) {
	api := s.App().Group("/api")
	v1 := api.Group("/v1")

	foods := v1.Group("/foods")
	foods.GET("", c.GetAll)
	foods.GET("/:id", c.GetDetail)
	foods.POST("", c.Create)
	foods.PATCH("/:id", c.Update)
	foods.DELETE("/:id", c.Delete)
}
