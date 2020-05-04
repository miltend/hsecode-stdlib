package matrix

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

import "github.com/cheekybits/genny/generic"

type ValueType generic.Type

type Matrix struct {
	matr []ValueType
	Rows int // number of rows
	Cols int // number of columns
}

func New(n, m int) *Matrix {
	if n <= 0 || m <= 0 {
		panic("rows and columns must be positive")
	}
	mtrx := Matrix{
		make([]ValueType, n*m),
		n,
		m,
	}

	return &mtrx
}

func (M *Matrix) Get(i, j int) ValueType {
	if i >= M.Rows || i < 0 || j >= M.Cols || j < 0 {
		panic("index is out of range")
	}
	return M.matr[M.Cols*i+j]
}

func (M *Matrix) Set(i, j int, v ValueType) {
	if i >= M.Rows || i < 0 || j >= M.Cols || j < 0 {
		panic("index is out of range")
	}
	M.matr[M.Cols*i+j] = v
}
