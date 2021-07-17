package main

import (
	"fmt"
	"time"
)

var nome = "Amabilie"
var momento = time.Now()
var horario = 10

func main() {
	Cumprimente(nome)
}

func Cumprimente(nome string) {
	switch {
	case horario > 0 && horario < 6:
		fmt.Printf("Olá, %s, agora é %d horas e é madrugada", nome, horario)
	case horario < 12:
		fmt.Printf("Olá, %s, agora é %d horas e é manhã", nome, horario)
	case horario >= 12 && horario < 18:
		fmt.Printf("Olá, %s,agora é %d horas e é tarde", nome, horario)
	case horario >= 18 && horario > 24:
		fmt.Printf("Olá, %s,agora é %d horas e é noite", nome, horario)
	default:
		fmt.Printf("horario nao é válido")
	}

}
