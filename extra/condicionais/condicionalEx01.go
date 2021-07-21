package main

import "fmt"

func main() {
	valor1 := 200
	valor2 := 300
	valor3 := 400
	maiorValor := 0

	if valor1 > valor2 && valor1 > valor3 {
		maiorValor = valor1
	} else if valor2 > valor1 && valor2 > valor3 {
		maiorValor = valor2
	} else {
		maiorValor = valor3
	}

	fmt.Println(maiorValor)
}
