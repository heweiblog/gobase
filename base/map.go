package main

import (
	"fmt"
	"sync"
)

//可以使用 make()，但不能使用 new() 来构造 map
func test1() {
	//m := make(map[string]uint8) //或下列
	m := map[string]uint8{}
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3
	m["4"] = 4
	fmt.Println(m, len(m))

	v := m["5"]
	fmt.Println(v)

	//检测map中是否存在零值时，不能通过值来判断，因为不存在会返回零值，
	//正确的做法是通过访问返回的第二个参数来判断
	if v, ok := m["5"]; ok {
		fmt.Println("in map", v, ok)
	} else {
		fmt.Println("not in map", ok)
	}

	//map 打印结果是无序的
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k, _ := range m {
		fmt.Println(k)
	}

	// 删除元素 删除不存在的元素不会报错
	//delete(m, "2")
	delete(m, "6")
	fmt.Println(m, len(m))

	//Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
	m = make(map[string]uint8)
	fmt.Println(m, len(m))

}

//Range() 方法可以遍历 sync.Map，遍历需要提供一个匿名函数，参数为 k、v，类型为 interface{}，每次 Range() 在遍历一个元素时，都会调用这个匿名函数把结果返回
//sync.Map 没有提供获取 map 数量的方法，替代方法是在获取 sync.Map 时遍历自行计算数量，sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
func iter(k, v interface{}) bool {
	fmt.Println("iterate:", k, v)
	if v == "exit" {
		return false
	}
	return true
}

// 这个就厉害了 就很像python的字典了
//Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的 可使用效率较高的并发安全的 sync.Map
func test2() {
	//sync.Map 不能使用 make 创建。
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", "xiaoxiao")
	//sync.Map 将键和值以 interface{} 类型进行保存。
	scene.Store("london", []string{"xxx", "sss"})
	scene.Store("egypt", 200)
	scene.Store(2, "exit")
	scene.Store(5, 11.11)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	//sync.Map 的 Delete 可以使用指定的键将对应的键值对删除。
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(iter)
}

func main() {
	test2()
}
