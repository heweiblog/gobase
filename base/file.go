package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// open 测试操作一个文件
func test() {
	//定义所有变量
	var (
		fp       *os.File
		err      error
		num      int
		readByte []byte
	)

	//判断打开文件
	if fp, err = os.Open("bit.go"); err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	//读文件
	readByte = make([]byte, 1024)
	if num, err = fp.Read(readByte); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num, string(readByte))
}

// ioutil ioutil读文件更高效
func ioutilRead() {
	content, err := ioutil.ReadFile("./bit.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	test()
	ioutilRead()
}
