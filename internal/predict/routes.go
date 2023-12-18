package predict

import "github.com/smartfood-capstone/backend/internal/server"

func RegisterRoute(s server.Server, c controller) {
	api := s.App().Group("/api")
	v1 := api.Group("/v1")

	predict := v1.Group("/predict")
	predict.GET("", c.DetectFood)
}
