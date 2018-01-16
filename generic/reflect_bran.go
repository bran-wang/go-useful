package main

import (
	"fmt"
	"reflect"
)

func index(slice interface{}, v interface{}) int {
	if slice := reflect.ValueOf(slice); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(v, slice.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}

func main() {
	slice1 := []int{0, 4, 5, 7, 23, 2}
	fmt.Println(index(slice1, 7))

	slice2 := []float64{3.3, 6.5, 4.8, 9.3, 16.5}
	fmt.Println(index(slice2, 6.5))

	slice3 := []interface{}{1, 2, 5, 4.3, "asd", 9}
	fmt.Println(index(slice3, 4))
}
