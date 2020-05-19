package ndarray

type NDArray struct {
	shape []int
}

func New(shape ...int) *NDArray {
	for i := 0; i < len(shape); i++ {
		if shape[i] <= 0 {
			panic("array size must be positive")
		}
	}
	arr := NDArray{shape: shape}

	return &arr
}

func (nda *NDArray) Idx(indices ...int) int {
	for i := 0; i < len(indices); i++ {
		if indices[i] < 0 || indices[i] >= nda.shape[i] {
			panic("index must not be negative or greater than array size")
		}
	}
	sum := 0
	for k := 1; k <= len(nda.shape); k++ {
		mult := 1
		for l := k + 1; l <= len(nda.shape); l++ {
			mult = mult * nda.shape[l-1]
		}
		sum = sum + mult*indices[k-1]
	}

	return sum
}
