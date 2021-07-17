package main

import (
	"fmt"
)

func main() {

	//semelhante ao foreach
	var meuSliceParaFor = []string{"for com contador", "for condicional", "for em algo", "E acabou a forezada!"}

	for indice, valor := range meuSliceParaFor {
		fmt.Printf("%d - %s\n", indice, valor)
	}

	for indice := range meuSliceParaFor {
		fmt.Printf("%d", indice)
	}

	fmt.Println("outro exemplo:")
	var fatia = []string{"zero", "um", "dois", "tres", "quatro"}
	for indice, valor := range fatia {
		fmt.Println(indice, valor)
	}
}
