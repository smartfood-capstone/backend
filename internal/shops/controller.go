package shops

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
	Name      string  `json:"name" validate:"required,printascii"`
	Location  string  `json:"location" validate:"required,printascii"`
	GmapsLink string  `json:"gmaps_link" validate:"required,url"`
	Latitude  float64 `json:"latitude" validate:"required,number"`
	Longitude float64 `json:"longitude" validate:"required,number"`
	Image     string  `json:"image" validate:"required,url"`
}

type requestUpdate struct {
	Name      string  `json:"name" validate:"omitempty,printascii"`
	Location  string  `json:"location" validate:"omitempty,printascii"`
	GmapsLink string  `json:"gmaps_link" validate:"omitempty,url"`
	Latitude  float64 `json:"latitude" validate:"omitempty,number"`
	Longitude float64 `json:"longitude" validate:"omitempty,number"`
	Image     string  `json:"image" validate:"omitempty,url"`
}

func (ru *requestUpdate) checkEmpty() bool {
	return ru.Name == "" && ru.Location == "" && ru.GmapsLink == "" && ru.Latitude == 0 && ru.Longitude == 0 && ru.Image == ""
}

func (c *controller) GetAll(ctx echo.Context) error {
	// mockResponse := `[
	//   {
	//     "id": 1,
	//     "name": "Shop 1",
	//     "location": "Shop 1 Location",
	//     "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
	//   },
	//   {
	//     "id": 2,
	//     "name": "Shop 1",
	//     "location": "Shop 2 Location",
	//     "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
	//   },
	//   {
	//     "id": 3,
	//     "name": "Shop 1",
	//     "location": "Shop 1 Location",
	//     "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
	//   },
	//   {
	//     "id": 4,
	//     "name": "Shop 1",
	//     "location": "Shop 2 Location",
	//     "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
	//   },
	//   {
	//     "id": 5,
	//     "name": "Shop 1",
	//     "location": "Shop 1 Location",
	//     "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
	//   },
	//   {
	//     "id": 6,
	//     "name": "Shop 1",
	//     "location": "Shop 2 Location",
	//     "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s"
	//   }
	// ]`
	// var resp any
	// json.Unmarshal([]byte(mockResponse), &resp)

	// return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))

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
	// mockResponse := `{
	//   "id": 2,
	//   "name": "Shop 1",
	//   "location": "Shop 2 Location",
	//   "latitude": 1.0,
	//   "longitude": 1.0,
	//   "image": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQGifwJSVbCfdxshbcC8_eyCWbiXucVwjSUPiQg5Ldpxg&s",
	//   "shops": [
	//     {
	//       "id": 1,
	//       "name": "shop 1",
	//       "price": 1.0,
	//       "image": "https://buffer.com/cdn-cgi/image/w=1000,fit=contain,q=90,f=auto/library/content/images/size/w1200/2023/10/free-images.jpg"
	//     }
	//   ]
	// }`

	// var resp any
	// json.Unmarshal([]byte(mockResponse), &resp)

	// return ctx.JSON(200, util.MakeResponse(200, "OK", nil, resp))
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

	shop := Shop{
		Name:      req.Name,
		Location:  req.Location,
		GmapsLink: req.GmapsLink,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Image:     req.Image,
	}

	resp, err := c.s.Create(controllerContext, shop)
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

	shop := Shop{
		Name:      req.Name,
		Location:  req.Location,
		GmapsLink: req.GmapsLink,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Image:     req.Image,
	}

	resp, err := c.s.Update(controllerContext, shop, id)
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
