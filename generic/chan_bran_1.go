package main

import (
	"fmt"
)

func print(ch chan int) {
	ch <- 1
	fmt.Println("hello world")
	//ch <- 1
}

func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go print(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

/*
OUTPUT:
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
*/
