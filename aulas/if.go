package main

import (
	"fmt"
)

func main() {
	idade := 65
	condicao := idade > 60
	fmt.Println(condicao)
	if condicao {
		fmt.Println("A pessoa é maior de 60 anos de idade")
	} else {
		fmt.Println("idade é abaixo de 60")
	}

	//melhor modo, n usar else, e sim um return, conforme ex abaixo:

	dia := "sabado"

	if dia == "sabado" {
		fmt.Println("O dia eh sabado")
		return
	}
	fmt.Println("o dia nao eh sabado")
}
