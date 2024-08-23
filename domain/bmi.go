package domain

type BMIRequest struct {
	Name string `json:"name"`
	Weight float32 `json:"weight" validate:"required"`
	Height float32 `json:"height" validate:"required"`
}

type BMIResponse struct {
	Bmi float32 `json:"bmi"`
}

