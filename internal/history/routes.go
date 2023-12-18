package history

import "github.com/smartfood-capstone/backend/internal/server"

func RegisterRoute(s server.Server, c IController) {
	api := s.App().Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/users")
	user.GET("/:id/history", c.GetUserHistory)
}
