package main

import (
	"fmt"
)

//递归算法
func JieCheng(n uint) uint {
	if n == 1 {
		return 1
	}
	return n * JieCheng(n-1)
}

//强制转换类型 类型(变量)
func main() {
	var m float32 = 5.0
	var n uint = uint(m)
	fmt.Println(JieCheng(n))
}
