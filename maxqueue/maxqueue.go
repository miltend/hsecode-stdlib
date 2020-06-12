package maxqueue

//go:generate genny -in=$GOFILE -out=int/dont_edit.go gen "ValueType=int"

import "github.com/cheekybits/genny/generic"

type ValueType generic.Type

type MaxQueue struct {
	// contains filtered or unexported fields
}

func New() *MaxQueue {

}
