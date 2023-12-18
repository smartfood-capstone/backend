package history

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	GetUserHistory(ctx echo.Context) error
}

type controller struct {
	l *logrus.Logger
}

func NewController(l *logrus.Logger) IController {
	return &controller{
		l: l,
	}
}

func (c *controller) GetUserHistory(ctx echo.Context) error {
	mockResponse := `[
    {
      "id": 1,
      "name": "Burger",
      "category": "bakso",
      "created_at": "2023-01-01 10:00:00",
      "description": "A burger",
      "image": "google.com"
    },
    {
      "id": 2,
      "name": "Burger",
      "category": "bakso",
      "created_at": "2023-01-01 11:00:00",
      "description": "A burger",
      "image": "google.com"
    }
  ]
`
	var resp any
	json.Unmarshal([]byte(mockResponse), &resp)
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}
