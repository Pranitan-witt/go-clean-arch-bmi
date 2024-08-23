package bmi

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain"
)

type Service struct {
	repo BMIRepositoryInterface
}

type BMIRepositoryInterface interface {
	SaveBMI(ctx context.Context, name string, bmi float32) error
}

func NewBMIService(r BMIRepositoryInterface) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CalcBMI(ctx context.Context, req *domain.BMIRequest) (*domain.BMIResponse, error) {
	bmi, err := s.CalculateBMI(req.Height, req.Weight)
	if err != nil {
		return nil, fmt.Errorf("calculateBMI| %s", err.Error())
	}

	

	if err := s.repo.SaveBMI(ctx, req.Name, bmi); err != nil {
		return nil, fmt.Errorf("repo.CalculateBMI| %s", err.Error())
	}

	return &domain.BMIResponse{
		Bmi:   bmi,
	},nil

}

func (s *Service) CalculateBMI(height, weight float32) (float32, error) {

	cal := weight / (height * height)
	if cal < 0 {
		return 0, fmt.Errorf("convertDecimalPlace| invalid BMI: %f", cal)
	}

	r, err := s.convertDecimalPlace(fmt.Sprintf("%.2f", cal))
	if err != nil {
		return 0, fmt.Errorf("convertDecimalPlace| %s", err.Error())
	}

	return r, nil
}

// convert 2 deimal places
func (s *Service) convertDecimalPlace(valueFormat string) (float32, error) {
	r, err := strconv.ParseFloat(valueFormat, 32)
	if err != nil {
		return 0, fmt.Errorf("parseFloat: %s", err.Error())
	}
	return float32(r), nil
}


