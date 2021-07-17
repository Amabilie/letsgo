package main

import "fmt"

//struct é tipo um "objeto"

type Pessoa struct {
	nome        string
	idade       int
	corPreferia string
}

func main() {
	pessoa1 := Pessoa{
		nome:        "Amabilie",
		idade:       29,
		corPreferia: "vermelho",
	}
	pessoa2 := Pessoa{
		nome:        "Marlon",
		idade:       29,
		corPreferia: "azul",
	}
	pessoa3 := Pessoa{
		nome:        "Joana",
		idade:       29,
		corPreferia: "rosa",
	}

	fmt.Println("As pessoas são", pessoa1.nome, pessoa2.nome, pessoa3.nome)

}
