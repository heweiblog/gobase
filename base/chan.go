package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/*
关于通道的一些知识点
1-内部实现同步，多个goruntine并发安全
2-通道默认为同步模式，通过建立有缓冲通道实现异步
*/

type Sync struct {
	w *sync.WaitGroup
	//m *sync.Mutex
	n map[uint64]int
}

//获取协程id
func GetId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func fetch(s *Sync, c chan string) {
	defer s.w.Done()
	num := 0
restart:
	select {
	case url := <-c:
		_, err := http.Get("http://" + url)
		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println(url)
			num++
			goto restart
		}
	//超时
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	}
	s.n[GetId()] = num
}

func test1() {
	//runtime.GOMAXPROCS(4)

	s := new(Sync)
	s.w = new(sync.WaitGroup)
	s.n = make(map[uint64]int, 100)

	ch := make(chan string, 100)

	start := time.Now()
	urls := []string{"www.qq.com", "www.baidu.com", "www.163.com"}

	for i := 0; i < 5; i++ {
		s.w.Add(1)
		go fetch(s, ch)
	}

	fmt.Println("send start")
	for i := 0; i < 10; i++ {
		for _, v := range urls {
			ch <- v
		}
	}
	fmt.Println("send end")

	s.w.Wait()
	fmt.Println(s)

	end := time.Now()

	fmt.Println(end.Sub(start))

}

//向已经关闭的channel发送数据会引发panic
func test2() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get the first result
	fmt.Println(<-ch)
	close(ch) //not ok (you still have other senders)
	//do other work
	time.Sleep(2 * time.Second)
}

func test3() {
	data := make(chan int)  // 数据交换队列
	exit := make(chan bool) // 退出通知
	go func() {
		for d := range data { // 从队列迭代接收数据，直到 close 。
			fmt.Println(d)
		}
		<-data
		fmt.Println("recv over.")
		exit <- true // 发出退出通知。
	}()
	data <- 1 // 发送数据。
	data <- 2
	data <- 3
	close(data) // 关闭队列。
	fmt.Println("send over.")
	<-exit // 等待退出通知。
}

//单向通道使用举例
func test4() {
	ch := make(chan int)
	exit := make(chan bool)
	var recv <-chan int = ch
	var send chan<- int = ch
	go func() {
		for d := range recv {
			fmt.Println(d)
		}
		exit <- true
	}()
	send <- 1
	//close(send) // 关闭队列。
	close(ch) // 关闭队列。
	fmt.Println("send over.")
	res := <-exit // 等待退出通知。
	fmt.Println("exit:", res)

}

func main() {
	test4()
}
