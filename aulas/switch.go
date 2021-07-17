package main

import (
	"fmt"
)

func main() {
	idade := 25

	switch {
	case idade >= 60:
		fmt.Printf("a pessoa tem idade %d e é idosa", idade)
	case idade >= 18 && idade < 60:
		fmt.Printf("a pessoa tem idade %d e é maior de idade", idade)
	case idade < 0:
		fmt.Printf("a pessoa tem idade %d e é inválida", idade)
	default:
		fmt.Printf("nenhuma situacao atendida")
	}
}
