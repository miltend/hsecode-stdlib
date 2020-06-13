//package tree
package main

import (
	"fmt"
	"strconv"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func EncodeTree(T *Tree) []string {
	var nodes []string
	var smolNodes []string

	if T.Left != nil {
		nodes = append(nodes, strconv.Itoa(T.Left.Value))
		smolNodes = append(smolNodes, EncodeTree(T.Left)...)
	} else {
		nodes = append(nodes, "nil")
		smolNodes = append(smolNodes, "nil", "nil")
	}
	if T.Right != nil {
		nodes = append(nodes, strconv.Itoa(T.Right.Value))
		smolNodes = append(smolNodes, EncodeTree(T.Right)...)
	} else {
		nodes = append(nodes, "nil")
		smolNodes = append(smolNodes, "nil", "nil")
	}
	nodes = append(nodes, smolNodes...)
	return nodes
}

func (T *Tree) Encode() []string {
	if T == nil {
		return nil
	}
	var nodes []string
	nodes = append(nodes, strconv.Itoa(T.Value))
	nodes = append(nodes, EncodeTree(T)...)

	for nodes[len(nodes)-1] == "nil" {
		nodes = nodes[:len(nodes)-1]
	}
	return nodes
}

func (T *Tree) NoLeft() *Tree {

	if T.Left == nil && T.Right == nil {
		return T
	}
	if T.Left == nil && T.Right != nil {
		tree := T.Right.NoLeft()
		return tree
	}
	if T.Left != nil && T.Right == nil {

		//if T.Left.Right == nil {}

		empVal := &Tree{Value: T.Value}
		T.Value = T.Left.Value
		T.Right = empVal
		tree := T.Left.NoLeft()
		//T.Left = nil
		return tree
	}
	if T.Left != nil && T.Right != nil {
		empVal := &Tree{Value: T.Value}
		empVal2 := &Tree{Value: T.Right.Value}
		T.Value = T.Left.Value
		T.Right = empVal
		T.Right.Right = empVal2
		//T.Left = T.Left.Left
		tree := T.Left.NoLeft()
		tree = T.Right.NoLeft()
		//T.NoLeft()
		//T.Left = nil
		return tree
	}

	return nil
}

func main() {
	//var empTree Tree
	t := &Tree{Value: 6, Left: &Tree{Value: 2, Left: &Tree{Value: 1, Left: &Tree{Value: 0}, Right: &Tree{Value: 3}}, Right: &Tree{Value: 4}}, Right: &Tree{Value: 8}}
	//t:= &Tree{
	//	Value: 2,
	//	Left:  &Tree{Value:1},
	//	Right: &Tree{Value:3},
	//}
	//new_tree := t
	fmt.Println(t.Encode())
	new_tree := t.NoLeft()
	//fmt.Println(new_tree.Encode())
	fmt.Println(new_tree.Left)
}
