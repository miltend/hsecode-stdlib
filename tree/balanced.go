package tree

import (
	"sort"
)

func NewBST(elements []int) *Tree {

	unique := make([]int, 0)
	h := make(map[int]struct{})

	for _, e := range elements {
		if _, ok := h[e]; !ok {
			h[e] = struct{}{}
			unique = append(unique, e)
		}
	}

	sort.Ints(unique)

	return fromSorted(unique)
}

func fromSorted(unique []int) *Tree {
	var root int
	var left []int
	var right []int
	if len(unique)%2 == 0 {
		root = unique[len(unique)/2]
		left = unique[:len(unique)/2]
		right = unique[len(unique)/2+1:]
	} else {
		root = unique[(len(unique)-1)/2]
		left = unique[:(len(unique)-1)/2]
		right = unique[(len(unique)-1)/2+1:]
	}

	t := Tree{
		Value: root,
		Left:  fromSorted(left),
		Right: fromSorted(right),
	}

	return &t
}
