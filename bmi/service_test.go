package bmi_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/bxcodec/go-clean-arch/bmi/mocks"
	"github.com/bxcodec/go-clean-arch/domain"
)

func TestCalculate(t *testing.T) {
	mockBMIRepo := new(mocks.BMIRepositoryInterface)

	t.Run("success", func(t *testing.T)  {
		s := bmi.NewBMIService(mockBMIRepo)

		r, err := s.CalculateBMI(1.55, 45)

		var expected float32 = 18.73
		assert.NoError(t, err)
		assert.Equal(t, expected, r)
	})

	t.Run("bmi < 0",func(t *testing.T) {
		s := bmi.NewBMIService(mockBMIRepo)

		r, err := s.CalculateBMI(1.55, -1)


		var expected float32 = 0
		assert.ErrorContains(t, err, "invalid BMI")
		assert.Equal(t, expected, r)
	})

}

func TestCalculateWithSave(t *testing.T){

	mockBMIRepo := new(mocks.BMIRepositoryInterface)
	mockBMIReq := domain.BMIRequest{
		Name:   "eiei",
		Weight: 45,
		Height: 1.55,
	}

	t.Run("success", func(t *testing.T) {
		mockBMIRepo.On("SaveBMI", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("float32")).Return(nil)
		s := bmi.NewBMIService(mockBMIRepo)
		r, err := s.CalcBMI(context.TODO(), &mockBMIReq)
		expected := &domain.BMIResponse{
			Bmi: 18.73,
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, r)
	})
	t.Run("save-error", func(t *testing.T) {
		mockBMIRepo.On("SaveBMI", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("float32")).Return(errors.New("error to save bmi"))
		s := bmi.NewBMIService(mockBMIRepo)
		r, err := s.CalcBMI(context.TODO(), &mockBMIReq)
		
		var expected *domain.BMIResponse

		assert.Error(t, err)
		assert.Equal(t, expected, r)
	})
}