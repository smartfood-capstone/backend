package shops

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	GetAllShops(ctx echo.Context) error
	GetShopDetail(ctx echo.Context) error
}

type controller struct {
	l *logrus.Logger
}

func NewController(l *logrus.Logger) IController {
	return &controller{
		l: l,
	}
}

func (c *controller) GetAllShops(ctx echo.Context) error {
	mockResponse := `[
    {
      "id": 1,
      "name": "Shop 1",
      "location": "Shop 1 Location",
      "image": "https://placehold.co/600x400"
    },
    {
      "id": 2,
      "name": "Shop 1",
      "location": "Shop 2 Location",
      "image": "https://placehold.co/600x400"
    },
    {
      "id": 3,
      "name": "Shop 1",
      "location": "Shop 1 Location",
      "image": "https://placehold.co/600x400"
    },
    {
      "id": 4,
      "name": "Shop 1",
      "location": "Shop 2 Location",
      "image": "https://placehold.co/600x400"
    },
    {
      "id": 5,
      "name": "Shop 1",
      "location": "Shop 1 Location",
      "image": "https://placehold.co/600x400"
    },
    {
      "id": 6,
      "name": "Shop 1",
      "location": "Shop 2 Location",
      "image": "https://placehold.co/600x400"
    }
  ]`
	var resp any
	json.Unmarshal([]byte(mockResponse), &resp)

	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}

func (c *controller) GetShopDetail(ctx echo.Context) error {
	mockResponse := `{
    "id": 2,
    "name": "Shop 1",
    "location": "Shop 2 Location",
    "latitude": 1.0,
    "longitude": 1.0,
    "image": "https://placehold.co/600x400",
    "foods": [
      {
        "id": 1,
        "name": "Food 1",
        "price": 1.0,
        "image": "https://placehold.co/600x400"
      }
    ]
  }`

	var resp any
	json.Unmarshal([]byte(mockResponse), &resp)

	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}
