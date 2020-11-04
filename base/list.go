package main

import (
	"container/list"
	"fmt"
	"sync"
)

func test() {
	l := list.New()
	// 尾部添加
	l.PushBack("canon")
	// 头部添加
	l.PushFront(67)
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter([]string{"high", "low"}, element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	l.PushBack([]int{1, 2, 3})

	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	l.PushBack(m)

	var scene sync.Map
	scene.Store("greece", "xiaoxiao")
	scene.Store("london", []string{"xxx", "sss"})
	scene.Store("egypt", 200)

	l.PushBack(scene)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	fmt.Println("----------------")

	// 删除
	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func main() {
	test()
}
