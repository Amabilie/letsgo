package main

import (
	"fmt"
)

func main() {
	horario := 14

	switch {
	case horario > 0 && horario < 6:
		fmt.Printf("agora é %d horas e é madrugada", horario)
	case horario < 12:
		fmt.Printf("agora é %d horas e é manhã", horario)
	case horario >= 12 && horario < 18:
		fmt.Printf("agora é %d horas e é tarde", horario)
	case horario >= 18 && horario > 24:
		fmt.Printf("agora é %d horas e é noite", horario)
	default:
		fmt.Printf("horario nao é válido")
	}
}
