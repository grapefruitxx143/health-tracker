package models

const (
	MinHealthyBMI = 18.5
	MaxHealthyBMI = 24.9
)

type User struct {
	Name   string    `json:"name"`
	Age    int       `json:"age"`
	Gender string    `json:"gender"`
	Height float64   `json:"height"`
	Weight []float64 `json:"weight"`
}

func (u User) CurrentWeight() (float64, bool) {
	if len(u.Weight) == 0 {
		return 0, false
	}
	return u.Weight[len(u.Weight)-1], true
}

func (u User) CalculateBMI() (float64, bool) {
	var BMI float64
	hInMeter := u.Height / 100

	CurrWeight, ok := u.CurrentWeight()
	if !ok {
		return 0, false
	}

	BMI = CurrWeight / (hInMeter * hInMeter)
	return BMI, true
}

func (u User) GetWeightRange() (float64, float64) {
	hInMeter := u.Height / 100
	MinWeight := MinHealthyBMI * (hInMeter * hInMeter)
	MaxWeight := MaxHealthyBMI * (hInMeter * hInMeter)
	return MinWeight, MaxWeight
}

func (u User) CalculateBMR() (float64, bool) {

	CurrWeight, ok := u.CurrentWeight()

	if !ok {
		return 0, false
	}

	baseBMR := (10 * CurrWeight) + (6.25 * u.Height) - (5 * float64(u.Age))
	if u.Gender == "male" {
		return baseBMR + 5, true
	}
	return baseBMR - 161, true
}
