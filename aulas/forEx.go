package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 27; i++ {
		if i >= 13 && i <= 27 {
			fmt.Println(i)
		}
	}

	//outra forma
	fmt.Println("outra forma")

	x := 13
	for x <= 27 {
		fmt.Println(x)
		x++
	}
}
