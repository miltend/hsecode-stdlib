package main

import (
	"fmt"
	"strconv"
)

//type Tree struct {
//	Value int
//	Left  *Tree
//	Right *Tree
//}

//func EncodeTree (T *Tree) []string {
//	var nodes []string
//	var smolNodes []string
//
//	if T.Left != nil {
//		nodes = append(nodes, strconv.Itoa(T.Left.Value))
//		smolNodes = append(smolNodes, EncodeTree(T.Left)...)
//	} else {
//		nodes = append(nodes, "nil")
//		smolNodes = append(smolNodes, "nil", "nil")
//	}
//	if T.Right != nil {
//		nodes = append(nodes, strconv.Itoa(T.Right.Value))
//		smolNodes = append(smolNodes, EncodeTree(T.Right)...)
//	} else {
//		nodes = append(nodes, "nil")
//		smolNodes = append(smolNodes, "nil", "nil")
//	}
//	nodes = append(nodes, smolNodes...)
//	return nodes
//}
//
//
//func (T *Tree) Encode() []string {
//	if T == nil {
//		return nil
//	}
//	var nodes []string
//	nodes = append(nodes, strconv.Itoa(T.Value))
//	nodes = append(nodes, EncodeTree(T)...)
//
//	for nodes[len(nodes) -1] == "nil" {
//		nodes = nodes[:len(nodes)-1]
//	}
//	return nodes
//}

func smolDecode(data []string) *Tree {
	var t Tree
	var value int

	if data[0] != "nil" {
		value, _ = strconv.Atoi(data[0])
		t.Left = &Tree{Value: value}
	}

	if data[1] != "nil" {
		value, _ = strconv.Atoi(data[1])
		t.Right = &Tree{Value: value}
	}
	return &t
}

func Decode(data []string) (*Tree, error) {
	for data[len(data)-1] == "nil" {
		data = data[:len(data)-1]
	}
	var t Tree

	//level := 0
	t.Value, _ = strconv.Atoi(data[0])

	return &t, nil
}

func main() {
	//t:= &Tree{Value:1, Right:&Tree{Value:3, Left:&Tree{Value:2}, Right:&Tree{Value:4}}}
	//list := t.Encode()
	//var tree *Tree
	t := &Tree{Value: 6, Left: &Tree{Value: 2, Left: &Tree{Value: 1, Left: &Tree{Value: 0}, Right: &Tree{Value: 3}}, Right: &Tree{Value: 4}}, Right: &Tree{Value: 8}}

	//list:= []string{"1","nil","3", "nil", "nil", "2", "4"}
	//tree, _ = Decode(list)
	a_list := t.Encode()
	fmt.Println(a_list)
}
