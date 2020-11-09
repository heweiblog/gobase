package main

import (
	"fmt"
	"strconv"
	"time"
)

func test1() {
	c1 := make(chan bool)
	c2 := make(chan string)

	go func() {
		<-c2
		c1 <- true
	}()

sel:
	for {
		select {
		case <-c1:
			fmt.Println("exit")
			break sel
		case c2 <- "hello":
			fmt.Println("hello")
		}
	}
}

//nice的用法
//两个队列 优先处理c1 c1没有任务时处理c2
func test2(c1 chan int, c2 chan string, exit chan bool) {
	for {
		select {
		case <-exit:
			fmt.Println("exit")
			return
		case d1 := <-c1:
			fmt.Println("recv int", d1)
		case d2 := <-c2:
		sel:
			for {
				select {
				case d1 := <-c1:
					fmt.Println("recv int", d1)
				default:
					//fmt.Println("default")
					break sel
				}
			}
			fmt.Println("recv string", d2)
		}
	}
}

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan string, 10)
	exit := make(chan bool)
	go test2(c1, c2, exit)

	for i := 1; i < 20; i++ {
		c1 <- i
		c2 <- strconv.Itoa(i)
	}
	time.Sleep(2 * time.Second)
	exit <- true
}
