package predict

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/smartfood-capstone/backend/internal/config"
)

type service struct {
	r IRepository
	l *logrus.Logger
}

type IService interface {
	DetectFoodUsingExternal(ctx echo.Context) (PredictResponse, error)
}

func NewService(r IRepository, l *logrus.Logger) IService {
	return &service{
		r: r,
		l: l,
	}
}

func (s *service) DetectFoodUsingExternal(ctx echo.Context) (PredictResponse, error) {
	file, fileHeader, err := ctx.Request().FormFile("file")
	if err != nil {
		return PredictResponse{}, err
	}

	predictType := strings.ToUpper(ctx.Request().FormValue("type"))

	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)
	fileWriter, err := multipartWriter.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		return PredictResponse{}, err
	}

	_, err = io.Copy(fileWriter, file)

	if err != nil {
		return PredictResponse{}, err
	}

	multipartWriter.WriteField("type", predictType)

	multipartWriter.Close()

	cfg := config.New()

	req, err := http.NewRequest("POST", cfg.MLHost+"/predict", &requestBody)
	if err != nil {
		return PredictResponse{}, err
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PredictResponse{}, err
	}

	defer resp.Body.Close()

	var respBody any
	err = json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		return PredictResponse{}, err
	}

	reqValue := reflect.ValueOf(respBody)
	Category := reqValue.MapIndex(reflect.ValueOf("data")).Interface().(string)

	history := History{
		// Hardcoded for now, no user login yet
		UserId: "1",
		Result: Result{
			Data: Category,
		},
	}

	err = s.r.InsertHistory(ctx.Request().Context(), history)
	if err != nil {
		return PredictResponse{}, err
	}

	food, err := s.r.GetFoodDetailByCategory(ctx.Request().Context(), Category)
	if err != nil {
		return PredictResponse{}, err
	}
	response := PredictResponse{
		Category:    food.Category,
		Name:        food.Name,
		Description: food.Description,
		Image:       food.Image,
	}

	return response, nil
}
