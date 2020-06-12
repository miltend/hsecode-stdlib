//package radix
package main

import "fmt"

//func Sort(data []uint64){
//	some_slice := []uint64
//	for _, v := range data{
//		b := v
//		for i:= 0 ; i < 8; i++{
//			suk := uint8(b)
//			b >>= 8
//		}
//	}
//
//}

func main() {

	//a:= uint64(18446744073709551615)
	//b:= a
	//for i:=0 ; i < 8; i++ {
	//	fmt.Println(uint8(b))
	//	b >>= 8
	//}
	const maxbit = -1 << 31
	//a:= [] int{2, 3, 9, 5, 3, 7, 1}
	//for i, v := range a{
	//	fmt.Println(i, v)
	////}
	//fmt.Println(a)
	//sort.Ints(a)
	//fmt.Println(a)
	fmt.Println(maxbit)

}
