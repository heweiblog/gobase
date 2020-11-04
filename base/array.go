package main

import "fmt"

func test_val(arr [3]int) {
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	fmt.Println(arr)
}

func test_ptr(arr *[3]int) {
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	fmt.Println(arr)
}

func test_slice(arr []int) {
	arr[0] = 111
	arr[1] = 222
	arr[2] = 333
	fmt.Println(arr)
}

func main() {
	//数组为值传递
	arr := [3]int{1, 2, 3}
	fmt.Println("beforce test_val", arr)
	test_val(arr)
	fmt.Println("after test_val", arr)

	//传递数组指针可达到修改的目的
	a := [3]int{1, 2, 3}
	fmt.Println("beforce test_ptr", a)
	test_ptr(&a)
	fmt.Println("after test_ptr")
	for i, v := range a {
		fmt.Println(i, v)
	}

	//slice传递参数
	ar := []int{1, 2, 3}
	fmt.Println("beforce test_slice", ar)
	test_slice(ar)
	fmt.Println("after test_slice", ar)

	//多维数组 初始化需要两步
	duowei := make([][]int, 5)
	for i, _ := range duowei {
		duowei[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			duowei[i][j] = j + 1
		}
	}
	duowei[2] = append(duowei[2], 999)
	fmt.Println(duowei)
}
