package main

import "fmt"

func main() {
	var numeros = [4]int64{1, 2, 3, 4}
	fmt.Println(numeros)
	fmt.Println(numeros[0])

	var diasDaSemana [7]string
	fmt.Println(diasDaSemana)

	diasDaSemana[0] = "Domingo"
	diasDaSemana[1] = "Segunda"
	diasDaSemana[2] = "Terca"
	diasDaSemana[3] = "Quarta"
	diasDaSemana[4] = "Quinta"
	diasDaSemana[5] = "Sexta"
	diasDaSemana[6] = "Sabado"
	fmt.Println(diasDaSemana)

	//Escreva um programa que contenha um array de strings
	//o valor de cada elemento deve ser o número do índice escrito por extenso
	//print na tela o tipo do seu array e valores de seus elementos

	var elementos = [5]string{"zero", "um", "dois", "tres", "quatro"}

	fmt.Printf("O tipo é %T e o índice é %v", elementos[0], elementos[0])
	fmt.Printf("O tipo é %T e o índice é %v", elementos[1], elementos[1])
	fmt.Printf("O tipo é %T e o índice é %v", elementos[2], elementos[2])
	fmt.Printf("O tipo é %T e o índice é %v", elementos[3], elementos[3])
	fmt.Printf("O tipo é %T e o índice é %v", elementos[4], elementos[4])

}
