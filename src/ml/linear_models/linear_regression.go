package linear_models

import (
	"../base/dataframe"
)

type LinearRegression struct {
	fitted      bool
	coefficient []float64
}

func init() {

}

func (lr *LinearRegression) Fit(samples DataFrame,
	features DataFrame) error {

	rows, cols = samples.Shape()

}

func Predict() {

}
