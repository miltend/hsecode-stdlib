package matrix

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

import "github.com/cheekybits/genny/generic"

type ValueType generic.Type

type Matrix struct {
	matr [][]ValueType
	Rows int // number of rows
	Cols int // number of columns
}

func New(n, m int) *Matrix {
	if n <= 0 || m <= 0 {
		panic("rows and columns must be positive")
	}
	var matrrr [][]ValueType
	mtrx := Matrix{
		matrrr,
		n,
		m,
	}
	for i := 0; i < n; i++ {
		row := make([]ValueType, m)
		mtrx.matr = append(mtrx.matr, row)
	}
	return &mtrx
}

func (M *Matrix) Get(i, j int) ValueType {
	if i >= M.Rows || i < 0 || j >= M.Cols || j < 0 {
		panic("index is out of range")
	}
	return M.matr[i][j]
}

func (M *Matrix) Set(i, j int, v ValueType) {
	if i >= M.Rows || i < 0 || j >= M.Cols || j < 0 {
		panic("index is out of range")
	}
	M.matr[i][j] = v
}
