package go_math

import "math"

func Floor(value float64, nbDigits float64) float64 {
	pow := math.Pow(10., nbDigits)
	return math.Floor(value*pow) / pow
}

func Round(value float64, nbDigits float64) float64 {
	pow := math.Pow(10., nbDigits)
	return math.Round(value*pow) / pow
}

func Ceil(value float64, nbDigits float64) float64 {
	pow := math.Pow(10., nbDigits)
	return math.Ceil(value*pow) / pow
}
