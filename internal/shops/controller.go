package shops

import (
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
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) GetShopDetail(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}
