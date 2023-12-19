package shops

import "github.com/smartfood-capstone/backend/internal/server"

func RegisterRoute(s server.Server, c IController) {
	api := s.App().Group("/api")
	v1 := api.Group("/v1")

	shops := v1.Group("/shops")
	shops.GET("", c.GetAll)
	shops.GET("/:id", c.GetDetail)
	shops.POST("", c.Create)
	shops.PATCH("/:id", c.Update)
	shops.DELETE("/:id", c.Delete)
}
