package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "hello how old are you how do you do yes i do"
	words := strings.Split(word, " ")
	fmt.Println(words)
	m := make(map[string]int)
	for _, v := range words {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	fmt.Println(m)
}
