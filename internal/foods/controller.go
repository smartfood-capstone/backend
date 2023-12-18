package foods

import (
	"github.com/labstack/echo/v4"
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
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) GetFoodDetail(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}
