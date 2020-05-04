package matrix

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

import "github.com/cheekybits/genny/generic"

type ValueType generic.Type

type Matrix struct {
	Rows int // number of rows
	Cols int // number of columns
	// contains filtered or unexported fields
	matr [][]ValueType
}

func New(n, m int) *Matrix {
	if n <= 0 || m <= 0 {
		panic("rows and columns must be positive")
	}
	var matrrr [][]ValueType
	mtrx := Matrix{
		n,
		m,
		matrrr,
	}
	for i := 0; i < mtrx.Rows; i++ {
		row := make([]ValueType, mtrx.Cols)
		mtrx.matr = append(mtrx.matr, row)
	}
	smth := &mtrx
	return smth
}

func (M *Matrix) Set(i, j int, v ValueType) {
	if i >= M.Rows || i < 0 {
		panic("row indice is out of range")
	}
	if j >= M.Cols || j < 0 {
		panic("column indice is out of range")
	}
	M.matr[i][j] = v
}

func (M *Matrix) Get(i, j int) ValueType {
	if i >= M.Rows || i < 0 {
		panic("row indice is out of range")
	}
	if j >= M.Cols || j < 0 {
		panic("column indice is out of range")
	}
	return M.matr[i][j]
}
