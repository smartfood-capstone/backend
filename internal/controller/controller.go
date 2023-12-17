package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	HealthCheck(ctx echo.Context) error
	GetAllFoods(ctx echo.Context) error
	GetFoodDetail(ctx echo.Context) error
	GetAllShops(ctx echo.Context) error
	GetShopDetail(ctx echo.Context) error
	GetUserHistory(ctx echo.Context) error
	DetectFood(ctx echo.Context) error
}

type controller struct {
}

func New() IController {
	return &controller{}
}

func (c *controller) HealthCheck(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

// TODO: Implement all functions below

func (c *controller) GetAllFoods(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) GetFoodDetail(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) GetAllShops(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) GetShopDetail(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) GetUserHistory(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}

func (c *controller) DetectFood(ctx echo.Context) error {
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, nil))
}
