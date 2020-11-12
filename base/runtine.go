package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type Sync struct {
	w *sync.WaitGroup
	m *sync.Mutex
	n int
}

//go 中主进程并不会等待所有的go协程结束才结束，为了实现同步可使用sync.WaitGroup
func fetch(s *Sync, url string) {
	defer s.w.Done()
	_, err := http.Get("http://" + url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(res)
	//fmt.Println(url, "Done")
	s.m.Lock()
	s.n++
	s.m.Unlock()
}

//统计请求成功次数
func test1() {
	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	fmt.Println("core=", *numCores)
	runtime.GOMAXPROCS(4)
	s := new(Sync)
	fmt.Println(s)
	s.w = new(sync.WaitGroup)
	s.m = new(sync.Mutex)
	fmt.Println(s)

	start := time.Now()
	urls := []string{"www.qq.com", "www.baidu.com", "www.163.com"}
	for i := 0; i < 100; i++ {
		for _, v := range urls {
			s.w.Add(1)
			go fetch(s, v)
		}
	}
	s.w.Wait()
	fmt.Println("all Done")
	end := time.Now()

	fmt.Println(end.Sub(start))
	fmt.Println(s)
}

func test2() {
	// setting GOMAXPROCS to 2 gives +- 22% performance increase,
	// but increasing the number doesn't increase the performance
	// without GOMAXPROCS: +- 86000
	// setting GOMAXPROCS to 2: +- 105000
	// setting GOMAXPROCS to 3: +- 94000
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for i := 0; ; i++ {
		select {
		case v := <-ch1:
			fmt.Printf("%d - Received on channel 1: %d\n", i, v)
		case v := <-ch2:
			fmt.Printf("%d - Received on channel 2: %d\n", i, v)
		}
	}
}

func test3() {
	var numCores = flag.Int("n", 2, "number of CPU cores to use")
	fmt.Println("core=", *numCores)
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)
	slice := make([]int, 0xffff)
	for i := 0; i < 0xffff; i++ {
		//fmt.Println(i)
		slice[i] = i
	}
	fmt.Println("len=", len(slice))
}

func test4() {
	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
}

func main() {
	//test1()
	//test2()
	test4()
}
