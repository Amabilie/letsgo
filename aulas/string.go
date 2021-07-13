package main

import "fmt"

func main() {
	var nome = "Amabi"
	var frase string
	frase = "Estudar é muito bom"

	fmt.Print(frase)
	//%T diz o tipo da variavel
	fmt.Printf("\n%T, %T", nome, frase)
	//%v diz o valor
	fmt.Printf("\n%V, %V", nome, frase)

	// aspas duplas, interpre o %. Aspas simples, não interpreta

	// := chamam de marmota

	// %s especificador

	fmt.Printf("\n")

	var meuNome = "Amabilie Macedo"
	corPreferida := "vermelho"

	fmt.Printf("Meu nome é %s, \nminha cor preferida é %s.", meuNome, corPreferida)

}
