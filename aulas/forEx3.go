package main

import (
	"fmt"
)

func main() {
	hora := 0
	for hora <= 24 {
		for minuto := 0; minuto < 60; minuto++ {
			for segundos := 0; segundos < 60; segundos = segundos + 10 {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundos)
			}
		}
		hora++
	}
}
