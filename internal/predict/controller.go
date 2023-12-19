package predict

import (
	"strings"

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

	predictType := ctx.Request().FormValue("type")
	predictType = strings.ToUpper(predictType)
	if predictType != "JAJANAN" && predictType != "MAKANAN" {
		// For now we will just assume that the user wants to predict food if the type is not specified
		predictType = "MAKANAN"
	}

	defer file.Close()

	resp, err := c.s.DetectFoodUsingExternal(ctx)
	if err != nil {
		return ctx.JSON(500, util.MakeResponse(500, "Internal Server Error", err, nil))
	}

	return ctx.JSON(200, util.MakeResponse(200, "Predict success", nil, resp))
}
