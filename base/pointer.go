package main

import "fmt"

func main() {
	a := 10
	b := &a
	c := &b
	d := &c
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Printf("c:%p type:%T\n", c, c) // b:0xc00001a078 type:*int
	fmt.Printf("d:%p type:%T\n", d, d) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
	fmt.Println(*b)                    // 0xc00000e018
	fmt.Println(**c)                   // 0xc00000e018
	fmt.Println(*d)                    // 0xc00000e018
	**c = 100
	fmt.Println(a) // 0xc00000e018
}
