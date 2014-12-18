package base

import (
	"Math"
	"base/"
)

type DataType float64

//- ----------------------- Data Frame
type DataFrame struct {
	Row int
	Col int
	df  [][]DataType
}

func (df *DataFrame) Shape() (int, int) {
	return df.Row, df.Col
}

// ------------------------- Data Volume
type DataVol struct {
	sx int

	sy int

	depth int

	w []float64

	dw []float64
}

func (dataVol *DataVol) init(sx int, sy int, depth int) {

	if sx == 1 {

		dataVol.sx = 1
		dataVol.sy = 1
		dataVol.depth = depth

		var wArray, dwArray [depth]float64
		dataVol.w = wArray
		dataVol.dw = dwArray
		for i := 0; i < dataVol.depth; i++ {

			dataVol.w[i] = sx[i]
		}
	} else {

		dataVol.sx = sx
		dataVol.sy = sy
		dataVol.depth = depth
		var n = sx * sy * depth
		var wArray, dwArray [n]float64 // TODO Optimization
		dataVol.w = wArray
		dataVol.dw = dwArray

		var scale = Math.Sart(1.0 / n)
		for i := 0; i < n; i++ {

			dataVol.w[i] = GetRandn(0.0, scale)
		}
	}

}
