package predict

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/util"
)

type IController interface {
	DetectFood(ctx echo.Context) error
}

type controller struct {
	l *logrus.Logger
	s IService
}

func NewController(s IService, l *logrus.Logger) IController {
	return &controller{
		s: s,
		l: l,
	}
}

func (c *controller) DetectFood(ctx echo.Context) error {

	// Limit upload size to 32 MB
	err := ctx.Request().ParseMultipartForm(32 << 20)
	if err != nil {
		return ctx.JSON(400, util.MakeResponse(400, "Bad Request", err, nil))
	}

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return ctx.JSON(400, util.MakeResponse(400, "Bad Request", err, nil))
	}

	defer file.Close()

	resp, err := c.s.DetectFoodUsingExternal(ctx)
	if err != nil {
		return ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
	}

	// mockResponse := `{
	// "id": 1,
	// "name": "Burger",
	// "category": "bakso",
	// "created_at": "2023-01-01 10:00:00",
	// "description": "A burger",
	// "image": "https://placehold.co/600x400"}`
	// var resp any
	// json.Unmarshal([]byte(mockResponse), &resp)
	return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
}
