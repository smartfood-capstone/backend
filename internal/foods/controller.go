package foods

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	GetAll(ctx echo.Context) error
	GetDetail(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type controller struct {
	s IService
	l *logrus.Logger
}

func NewController(s IService, l *logrus.Logger) IController {
	return &controller{
		s: s,
		l: l,
	}
}

type requestCreate struct {
	Name        string `json:"name" validate:"required,printascii"`
	Description string `json:"description" validate:"required,printascii"`
	Category    string `json:"category" validate:"required,printascii"`
	Image       string `json:"image" validate:"required,url"`
}

type requestUpdate struct {
	Name        string `json:"name" validate:"omitempty,printascii"`
	Description string `json:"description" validate:"omitempty,printascii"`
	Category    string `json:"category" validate:"omitempty,printascii"`
	Image       string `json:"image" validate:"omitempty,url"`
}

func (ru *requestUpdate) checkEmpty() bool {
	return ru.Name == "" && ru.Description == "" && ru.Category == "" && ru.Image == ""
}

func (c *controller) GetAll(ctx echo.Context) error {
	controllerContext := ctx.Request().Context()

	qlimit := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		limit = 100
	}

	qpage := ctx.QueryParam("page")
	page, err := strconv.Atoi(qpage)
	if err != nil {
		page = 1
	}

	qname := ctx.QueryParam("name")

	if page <= 0 {
		page = 1
	}

	offset := limit * (page - 1)

	params := getAllRepoParams{
		Name:   qname,
		Limit:  limit,
		Offset: offset,
	}

	resp, err := c.s.GetAll(controllerContext, params)
	if err != nil {
		c.l.Errorf("error when getting data from service err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when fetching data", err, nil))
	}

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
}

func (c *controller) GetDetail(ctx echo.Context) error {
	controllerContext := ctx.Request().Context()

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.l.Errorf("error when converting id to integer id: %d, err: %s", id, err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "cannot parse id as integer", err, nil))
	}

	resp, err := c.s.GetDetail(controllerContext, id)
	if err != nil {
		c.l.Errorf("error when getting detail data from service id: %d, err: %s", id, err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when fetching data with current id", err, nil))
	}

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
}

func (c *controller) Create(ctx echo.Context) error {
	var req requestCreate
	controllerContext := ctx.Request().Context()

	if err := ctx.Bind(&req); err != nil {
		c.l.Errorf("invalid request body err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when parsing request body", err, nil))
	}

	v := validator.New()
	err := v.Struct(&req)
	if err != nil {
		c.l.Errorf("error when validating body err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when validating body", err, nil))
	}

	food := Food{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Image:       req.Image,
	}

	resp, err := c.s.Create(controllerContext, food)
	if err != nil {
		c.l.Errorf("error when creating data from service err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error creating data from service", err, nil))
	}

	return ctx.JSON(http.StatusCreated, util.MakeResponse(http.StatusCreated, "OK", nil, resp))
}

func (c *controller) Update(ctx echo.Context) error {
	var req requestUpdate
	controllerContext := ctx.Request().Context()

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.l.Errorf("error when converting id to integer id: %d, err: %s", id, err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "cannot parse id as integer", err, nil))
	}

	if err := ctx.Bind(&req); err != nil {
		c.l.Errorf("invalid request body err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when parsing request body", err, nil))
	}

	v := validator.New()
	err = v.Struct(&req)
	if err != nil {
		c.l.Errorf("error when validating body err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when validating body", err, nil))
	}

	if req.checkEmpty() {
		c.l.Error("request body cannot be empty")
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "request body cannot be empty", err, nil))

	}

	food := Food{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Image:       req.Image,
	}

	resp, err := c.s.Update(controllerContext, food, id)
	if err != nil {
		c.l.Errorf("error when update data from service id: %d, err: %s", id, err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error update data from service", err, nil))
	}

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
}

func (c *controller) Delete(ctx echo.Context) error {
	controllerContext := ctx.Request().Context()

	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.l.Errorf("error when converting id to integer id: %d, err: %s", id, err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "cannot parse id as integer", err, nil))
	}

	resp, err := c.s.Delete(controllerContext, id)
	if err != nil {
		c.l.Errorf("error when deleting data from service id: %d, err: %s", id, err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error deleting data from service", err, nil))
	}

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
}
