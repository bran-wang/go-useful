package main

import (
	"fmt"
)

type MyType interface {
	Equal(a int, b interface{}) bool
	Len() int
}

func index(slice MyType, b interface{}) int {
	for i := 0; i < slice.Len(); i++ {
		if slice.Equal(i, b) {
			return i
		}
	}
	return -1
}

type IntType []int

func (slice IntType) Equal(i int, b interface{}) bool {
	return slice[i] == b.(int)
}

func (slice IntType) Len() int {
	return len(slice)
}

func main() {
	slice1 := IntType{1, 2, 3, 4, 5}
	fmt.Printf("%d\n", index(slice1, 3))
}
