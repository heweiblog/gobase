package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
	// 声明一个等待组
	var wg sync.WaitGroup
	// 准备一系列的网站地址
	var urls = []string{
		"http://www.baidu.com/",
		"https://www.qq.com/",
		"https://www.jd.com/",
		"https://www.58.com/",
		"https://www.hao123.com/",
		"https://www.sina.com/",
		"https://www.taobao.com/",
		"https://www.sohu.com/",
	}
	start := time.Now()
	// 遍历这些地址
	for _, url := range urls {
		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)
		// 开启一个并发
		go func(url string) {
			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()
			// 使用http访问提供的地址
			_, err := http.Get(url)
			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(url, err)
			// 通过参数传递url地址
		}(url)
	}
	// 等待所有的任务完成
	wg.Wait()
	sub := time.Now().Sub(start)
	fmt.Println(sub)
	fmt.Println("over")
}
