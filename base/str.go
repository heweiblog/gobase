package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	s := "hello"
	//s[0] = 'H' 字符串不可修改
	fmt.Println(s[0])
	fmt.Println(s)
	bs := []byte(s)
	bs[0] = 'H'
	bs[1] = 'E'
	s = string(bs)
	//utf8.ValidString(s) 判断是否是utf8编码
	fmt.Println(s, utf8.ValidString(s), len(s))
	//这个在python长度为3 在go中字符串的长度为存储字符串字节的长度
	s1 := "He伟"
	s2 := "♥#&"
	fmt.Println(s1, utf8.ValidString(s1), len(s1), s2, utf8.ValidString(s2), len(s2))

	for i, v := range s1 {
		fmt.Println(i, v)
	}

	for i, v := range []byte(s1) {
		fmt.Println(i, v)
	}

	num := 0
	fmt.Println("num=", num)
	//++num  自增 自减 不支持前置操作，支持后置操作
	num++
	fmt.Println("num++  ->  ", num)
	fmt.Println("num=", num)
	num--
	fmt.Println("num--  ->  ", num)
	fmt.Println("num=", num)
	num += 3
	fmt.Println("num+=3  ->  ", num)

	// switch与C语言不同，不需要加入break，自动跳出，加入fallthrough关键字会执行下一个case
	switch 1 {
	case 1:
		log.Println("switch 1")
		fallthrough
	case 2:
		log.Println("switch 2")
	case 3:
		log.Println("switch 3")
	}

	log.Println("log println 1")
	//panic抛出异常后结束进程
	//log.Panicln("log fatal")
	//fatal直接结束进程
	log.Fatalln("log fatal")
	log.Println("log println 2")
}
