package tree

//package main

import (
	//"fmt"
	"sort"
)

//type Tree struct {
//	Value int
//	Left  *Tree
//	Right *Tree
//}

func NewBST(elements []int) *Tree {

	if len(elements) == 0 {
		return nil
	}
	unique := make([]int, 0)
	h := make(map[int]struct{})

	for _, e := range elements {
		if _, ok := h[e]; !ok {
			h[e] = struct{}{}
			unique = append(unique, e)
		}
	}

	sort.Ints(unique)
	//fmt.Println(unique)
	return fromSorted(unique)
}

func fromSorted(unique []int) *Tree {

	if len(unique) == 0 {
		return nil
	}
	var root int
	var left []int
	var right []int
	var t Tree

	if len(unique)%2 == 0 {
		root = unique[len(unique)/2]
		left = unique[:len(unique)/2]
		right = unique[len(unique)/2+1:]
	} else {
		root = unique[(len(unique)-1)/2]
		left = unique[:(len(unique)-1)/2]
		right = unique[(len(unique)-1)/2+1:]
	}
	//
	//fmt.Println(root)
	//fmt.Println(left)
	//fmt.Println(right)

	t = Tree{
		Value: root,
		Left:  fromSorted(left),
		Right: fromSorted(right),
	}

	return &t
}

//
//func main() {
//
//	//var a [] int
//	a:= [] int {1, 3, 4, 5}
//
//	tre := NewBST(a)
//	fmt.Println(tre)
//}
