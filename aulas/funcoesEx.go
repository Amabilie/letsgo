package main

import (
	"fmt"
)

var n1 = 2
var n2 = 3

func main() {
	soma := some(n1, n2)
	fmt.Println(soma)
}

func some(a, b int) int {
	soma := a + b
	return soma
}
