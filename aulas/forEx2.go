package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 24 {
		fmt.Printf("%.2d:00\n", i)
		i++
	}
}
