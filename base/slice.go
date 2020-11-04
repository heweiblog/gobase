package main

import (
	"fmt"
)

//copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。
// copy( destSlice, srcSlice []T) int
// 将 srcSlice 复制到 destSlice，目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。
func test1() {
	slice1 := []int{5, 4, 3, 2, 1}
	slice2 := []int{6, 7, 8}
	slice3 := []int{9, 10, 11}
	fmt.Println(slice1, "copy to", slice2)
	s1 := copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2, s1)
	fmt.Println(slice3, "copy to", slice1)
	s2 := copy(slice1, slice3) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1, s2)
}

func test2() {
	a := []int{5, 4, 3, 2, 1}

	//删除前两个
	a = a[2:]
	fmt.Println(a)

	// 删除后两个
	b := []int{5, 4, 3, 2, 1}
	b = b[:len(b)-1-2]
	fmt.Println(b)

	// 删除从第二个开始的3个
	c := []int{5, 4, 3, 2, 1}
	c = append(c[:2-1], c[2+3-1:]...)
	fmt.Println(c)
}

//向一个slice追加两一个slice后面必须有 ... append追加元素一定是从len开始追加
func test3() {
	a := []int{5, 4, 3, 2, 1}
	//在slice首部添加一个或多个元素
	a = append([]int{7, 6}, a...)
	fmt.Println(a)

	//在slice中间添加一个或多个元素
	//在第三个后面添加8和9两个元素
	a = append(a[:3], append([]int{8, 9}, a[3:]...)...)
	fmt.Println(a)
}

//range用法 range 返回的是每个元素的副本，而不是直接返回对该元素的引用
func test4() {
	// 创建一个整型切片，并赋值
	slice := []int{10, 20, 30, 40}
	fmt.Println(len(slice), cap(slice))
	// 迭代每个元素，并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice[index])
	}
}

func main() {
	test4()
}
