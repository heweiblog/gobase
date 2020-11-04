package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	c := 3
	d := -4
	e := 27
	f := 21

	//按位与 &
	fmt.Println(a & c)

	//按位或 |
	fmt.Println(a | b)

	//取反 ^ c语言中取反为~
	fmt.Println(^b)
	fmt.Println(^d + 1)

	//左移
	fmt.Println(a << 1)

	//右移
	fmt.Println(b >> 1)

	// 异或 c语言异或为^  相同位不管0或1均为假，e为真f为假才为真，e为假f为真为假
	// 简单来说即e为真f为假为真，其他均为假
	fmt.Println(e &^ f)
}
