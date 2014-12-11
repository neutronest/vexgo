package base

type DataType float64

type DataFrame struct {
	Row int
	Col int
	df  [][]DataType
}

func (df *DataFrame) Shape() (int, int) {
	return df.Row, df.Col
}
