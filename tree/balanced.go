package tree

//package main

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
	var t Tree

	if len(unique) == 0 {
		return &t
	}

	if len(unique)%2 == 0 {
		root = unique[len(unique)/2]
		left = unique[:len(unique)/2]
		right = unique[len(unique)/2+1:]
	} else {
		root = unique[(len(unique)-1)/2]
		left = unique[:(len(unique)-1)/2]
		right = unique[(len(unique)-1)/2+1:]
	}

	t = Tree{
		Value: root,
		Left:  fromSorted(left),
		Right: fromSorted(right),
	}

	return &t
}

//func main() {
//
//	var a [] int
//	a = nil
//	//root := a[len(a)/2]
//	//left := a[:len(a)/2]
//	//right := a[len(a)/2 + 1:]
//	root := a[(len(a) - 1) / 2]
//	left := a[:(len(a) - 1) / 2]
//	right := a[(len(a) - 1) / 2 + 1:]
//	fmt.Println(root)
//	fmt.Println(left)
//	fmt.Println(right)
//	//if len(a) == 0 {
//	//	fmt.Println("sadadsdas")
//	//}
//}
