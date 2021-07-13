package main

import (
	"fmt"
	"time"
)

func main() {
	var dataAtual = time.Now()
	var anoNascimento int
	fmt.Println("Digite o ano do seu nascimento: ")
	_, err := fmt.Scan(&anoNascimento)
	if err != nil {
		panic(err)
	}
	fmt.Println("a data atua é ", dataAtual)

	var idade = (dataAtual.Year() - anoNascimento)

	fmt.Println("Sua idade é ", idade)

}
