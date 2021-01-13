package main

import (
	"fmt"
)

func test() (int, int, string) {
	return 1, 100, "xxx"
}

func test1() (x int, y int, s string) {
	x = 111
	y = 222
	s = "aaa"
	return
}

type Func func() (int, int, string)

//test
func test2() {
	//x, y, s := test()
	//x, y, s := test1()
	//var f Func
	var f Func
	f = test1
	x, y, s := f()
	fmt.Println(x, y, s)

	c := make(chan bool)
	//匿名函数 一种方式是定义期间直接调用 另一种是将匿名函数赋值给变量后调用
	fu := func(num int) {
		fmt.Println("in go func num =", num)
		c <- true
	}
	go fu(10)
	fmt.Println(<-c)
}

func call(s []int, f func(int, int)) {
	for i, v := range s {
		f(i, v)
	}
}

//匿名函数用作回调函数 此设计很有创意
func test3() {
	a := []int{1, 2, 3, 4, 5, 6}
	f := func(a, b int) {
		fmt.Println(a, "*", b, "=", a*b)
	}
	call(a, f)
}

func main() {
	test3()
}
