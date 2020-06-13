package main

func (T *Tree) IsSym() bool {
	if T == nil {
		return true
	}
	if T.Left == nil && T.Right == nil {
		return true
	}
	if T.Left == nil || T.Right == nil {
		return false
	}
	return isAntySym(T.Left, T.Right)
}

func isAntySym(a, b *Tree) bool {
	var isAS bool
	if a.Value != b.Value {
		return isAS
	}

	if a.Right == nil && b.Left != nil {
		return isAS
	}

	if a.Left == nil && b.Right != nil {
		return isAS
	}

	if a.Left != nil && b.Right == nil {
		return isAS
	}
	if a.Right != nil && b.Left == nil {
		return isAS
	}

	if a.Right != nil && b.Left != nil {
		isAS = isAntySym(a.Right, b.Left)
		if !isAS {
			return isAS
		}
	}

	if a.Left != nil && b.Right != nil {
		isAS = isAntySym(a.Left, b.Right)
		if !isAS {
			return isAS
		}
	}

	isAS = true
	return isAS

}
