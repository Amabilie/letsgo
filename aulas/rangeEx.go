package main

import (
	"fmt"
)

func main() {

	listaDeMercado := []string{"abacate", "sabonete", "azeite", "pao", "sucrilhos"}
	for indice, valor := range listaDeMercado {
		fmt.Printf("%d) %s\n", indice, valor)
	}
}
