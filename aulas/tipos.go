package main

import (
	"fmt"
)

func main() {
	var nome = "Amabilie"
	idade := 29
	fmt.Println(nome, idade)
	fmt.Printf("o nome é %s e a idade é %d\n", nome, idade)
	fmt.Printf("O tipo de nome é %T e o tio de idade é %T", nome, idade)

	var names = [3]string{"Joana", "Maria", "Ana Julia"}
	fmt.Println(names)

	idades := make([]int, 3)
	idades[0] = 47
	idades[1] = 53
	idades[2] = 24
	fmt.Println(idades)
	idades = append(idades, 67, 55, 21) //é assim que aumenta o tamanho do array, usando o append
	fmt.Println(idades)

	var mapa = make(map[int]string // para declarar um map, primeiro o tipo da chave, depois o tipo do valor
	mapa["X"] = 10
	fmt.Println(mapa["X"])

}
