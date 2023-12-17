package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	HealthCheck(ctx echo.Context) error
}

type controller struct {
}

func New() IController {
	return &controller{}
}

func (c *controller) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}
