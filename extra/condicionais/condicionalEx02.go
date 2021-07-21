package main

import "fmt"

func main() {
	var numero int
	fmt.Println("Digite um numero intiro: ")
	_, err := fmt.Scan(&numero)
	if err != nil {
		panic(err)
	}

	resultadoCom2 := (numero % 2)
	resultadoCom3 := (numero % 3)

	if resultadoCom2ero / resultadoCom2 {
		fmt.Println("o número é ZERO")
	}
	fmt.Println(resultado)
}
