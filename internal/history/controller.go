package history

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	GetUserHistory(ctx echo.Context) error
}

type controller struct {
	l *logrus.Logger
	s IService
}

func NewController(l *logrus.Logger, s IService) IController {
	return &controller{
		l: l,
		s: s,
	}
}

func (c *controller) GetUserHistory(ctx echo.Context) error {
	controllerContext := ctx.Request().Context()

	qlimit := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(qlimit)
	if err != nil {
		limit = 25
	}

	qpage := ctx.QueryParam("page")
	page, err := strconv.Atoi(qpage)
	if err != nil {
		page = 1
	}

	if page <= 0 {
		page = 1
	}

	offset := limit * (page - 1)

	params := getAllParams{
		Limit:  limit,
		Offset: offset,
	}

	resp, err := c.s.GetAll(controllerContext, params)
	if err != nil {
		c.l.Errorf("error when getting data from service err: %s", err)
		return ctx.JSON(http.StatusBadRequest, util.MakeResponse(http.StatusBadRequest, "error when fetching data", err, nil))
	}

	return ctx.JSON(http.StatusOK, util.MakeResponse(http.StatusOK, "OK", nil, resp))
	//		mockResponse := `[
	//	    {
	//	      "id": 1,
	//	      "name": "Burger",
	//	      "category": "bakso",
	//	      "created_at": "2023-01-01 10:00:00",
	//	      "description": "A burger",
	//	      "image": "google.com"
	//	    },
	//	    {
	//	      "id": 2,
	//	      "name": "Burger",
	//	      "category": "bakso",
	//	      "created_at": "2023-01-01 11:00:00",
	//	      "description": "A burger",
	//	      "image": "google.com"
	//	    }
	//	  ]
	//
	// `
	//
	//	var resp any
	//	json.Unmarshal([]byte(mockResponse), &resp)
	//	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}
