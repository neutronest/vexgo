package base

import (
	"math"
	"math/rand"
	"time"
)

func GaussianRandom() {

	rand.Seed(99)

}

func GetRandf(a float64, b float64) float64 {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return (b-a)*r.Float64() + a // Float64 return 0~1
}

func GetRani(a float64, b float64) float64 {

	return math.Floor(GetRandf(a, b))
}

func GetRandn(mu float64, std float64) float64 {

	return rand.NormFloat64()*mu + std
}
