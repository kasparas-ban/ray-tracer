package utils

import "math/rand"

func RandomNum(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
