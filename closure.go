package main

import (
	"fmt"
	"runtime"
)

/*
闭包(closure)：是在其词法上下文中引用了自由变量的函数
我们可以观察到两个特征：
1. 返回的是函数，匿名函数
2. 函数里面有变量，这个变量将来会用到
*/

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s: %d\n", file, line)
	}

	where()

	where()

	//func2
	f := test(123)
	f()
}

/*
OUTPUT:
/Users/branw/Documents/code/go_moudle/go-useful/closure.go: 14
/Users/branw/Documents/code/go_moudle/go-useful/closure.go: 16
*/

func test(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

/*
OUTPUT:
123
*/
