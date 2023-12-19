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
      "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
    },
    {
      "id": 2,
      "name": "Shop 1",
      "location": "Shop 2 Location",
      "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
    },
    {
      "id": 3,
      "name": "Shop 1",
      "location": "Shop 1 Location",
      "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
    },
    {
      "id": 4,
      "name": "Shop 1",
      "location": "Shop 2 Location",
      "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
    },
    {
      "id": 5,
      "name": "Shop 1",
      "location": "Shop 1 Location",
      "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
    },
    {
      "id": 6,
      "name": "Shop 1",
      "location": "Shop 2 Location",
      "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
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
    "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s",
    "foods": [
      {
        "id": 1,
        "name": "Food 1",
        "price": 1.0,
        "image": "https://buffer.com/cdn-cgi/image/w=1000,fit=contain,q=90,f=auto/library/content/images/size/w1200/2023/10/free-images.jpg"
      }
    ]
  }`

	var resp any
	json.Unmarshal([]byte(mockResponse), &resp)

	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}
