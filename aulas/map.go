package main

import "fmt"

func main() {
	mapinha := make(map[string]int)

	mapinha["tres"] = 3
	mapinha["quatrocentos"] = 400
	fmt.Printf("%T, %V\n", mapinha, mapinha)
	fmt.Println(mapinha)

	//crie um map que as chaves sao os n meses, e o valor é mes representa
	//printe os meses janeiro e dezembro
	//printe o tamanho do map

	ano := make(map[int]string)
	ano[1] = "janeiro"
	ano[2] = "fevereiro"
	ano[3] = "marco"
	ano[4] = "abril"
	ano[5] = "maio"
	ano[6] = "junho"
	ano[7] = "julho"
	ano[8] = "agosto"
	ano[9] = "setembro"
	ano[10] = "outubro"
	ano[11] = "novembro"
	ano[12] = "dezembro"

	fmt.Println(ano)
	fmt.Printf("%v, %v\n", ano[1], ano[12])
	fmt.Println(len(ano))

	cores := map[string]string{
		"cinza": "#424242",
		"roxo":  "#a378f8",
	}

	cinza, existe := cores["azul"]
	fmt.Println(cinza)

	fmt.Println(existe)

	//deletar valor de um map

	fmt.Println("1", cores)
	delete(cores, "cinza")
	fmt.Println("2", cores)
}
