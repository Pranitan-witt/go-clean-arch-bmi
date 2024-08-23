package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bxcodec/go-clean-arch/domain"
)

type BMISerivce interface {
	CalcBMI(ctx context.Context, req *domain.BMIRequest) (*domain.BMIResponse, error)
	CalculateBMI(height, weight float32) (float32, error)
}

// ArticleHandler  represent the httphandler for article
type BMIHandler struct {
	Service BMISerivce
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewBMIHandler(e *echo.Echo, svc BMISerivce) {
	handler := &BMIHandler{
		Service: svc,
	}
	e.POST("/calculate-bmi", handler.CalculateBMI)
}

func (h *BMIHandler) CalculateBMI(c echo.Context) (err error) {
	var bmi domain.BMIRequest
	err = c.Bind(&bmi)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err = h.validate(bmi.Height, bmi.Weight); err != nil {
		return err
	}

	ctx := c.Request().Context()
	res, err := h.Service.CalcBMI(ctx, &bmi)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *BMIHandler) validate(height, weight float32) error {
	if weight <= 0 {
		return errors.New("weight should greater than 0")
	}

	if height > 3 {
		return errors.New("height should be in meter not centimeter")
	}

	return nil
}
