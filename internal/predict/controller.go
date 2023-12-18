package predict

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	DetectFood(ctx echo.Context) error
}

type controller struct {
	l *logrus.Logger
}

func NewController(l *logrus.Logger) IController {
	return &controller{
		l: l,
	}
}

func (c *controller) DetectFood(ctx echo.Context) error {
	mockResponse := `{
  "id": 1,
  "name": "Burger",
  "category": "bakso",
  "created_at": "2023-01-01 10:00:00",
  "description": "A burger",
  "image": "https://placehold.co/600x400"}`
	var resp any
	json.Unmarshal([]byte(mockResponse), &resp)
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}
