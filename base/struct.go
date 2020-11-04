package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//以小写字母开头的结构体将不会被（json、xml、gob等）编码，因此当你编码这些未导出的结构体时，你将会得到零值。
type People struct {
	Name string
	Age  uint8
}

type Stu struct {
	name string
	age  uint8
}

func main() {
	//var a People
	//a.Name = "贺伟伟"
	//a.Age = 18
	a := People{"hww", 19}
	fmt.Println(a)

	j, er := json.Marshal(a)
	if er != nil {
		log.Fatalln(er)
	}
	fmt.Println(string(j))

	var d People
	er = json.Unmarshal(j, &d)
	if er != nil {
		log.Fatalln(er)
	}
	fmt.Println(d)

	b := new(Stu)
	b.name = "mnn"
	b.age = 16
	fmt.Println(b)

	r, er := json.Marshal(b)
	if er != nil {
		log.Fatalln(er)
	}
	// Stu 中为小写 所以编码为默认零值
	fmt.Println(string(r))

	var s Stu
	er = json.Unmarshal(r, &s)
	if er != nil {
		log.Fatalln(er)
	}
	fmt.Println(s)
}
