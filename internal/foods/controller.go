package foods

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	GetAllFoods(ctx echo.Context) error
	GetFoodDetail(ctx echo.Context) error
}

type controller struct {
	l *logrus.Logger
}

func NewController(l *logrus.Logger) IController {
	return &controller{
		l: l,
	}
}

func (c *controller) GetAllFoods(ctx echo.Context) error {
	mockResponse := `[
    {"id": 1,
    "name": "Burger",
    "description": "A burger",
    "image": "https://placehold.co/600x400"
    },
    {
    "id": 2,
    "name": "Burger 2",
    "description": "A burger",
    "image": "https://placehold.co/600x400"}
  ]`

	var resp any
	json.Unmarshal([]byte(mockResponse), &resp)

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
}

func (c *controller) GetFoodDetail(ctx echo.Context) error {
	mockResponse :=
		`{
    "id": 1,
    "name": "Burger",
    "description": "A burger",
    "image": "https://placehold.co/600x400",
    "shops": [
      {
        "id": 1,
        "name": "McDonalds"
      },
      {
        "id": 2,
        "name": "Burger King"
      }
    ]
  }`

	var resp any
	err := json.Unmarshal([]byte(mockResponse), &resp)
	if err != nil {
		log.Error(err)
	}

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
}
