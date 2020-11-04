package main

import (
	"fmt"
)

func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
	InLoop:
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break InLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}
