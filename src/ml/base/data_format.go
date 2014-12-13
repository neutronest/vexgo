package base

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

	w float64

	dw float64
}

func (dataVol *DataVol) init(sx int, sy int, depth int, weight float64) {

	dataVol.sx = 1
	return nil
}
