package main

import "fmt"

func main() {
	var listaDeCompras = []string{"agua", "ovos", "saco de lixo"}

	fmt.Println(listaDeCompras)
	fmt.Println(listaDeCompras[0])

	var listaDaIrma = []string{"detergente", "chocolate"}

	//para add elementos, usa append, como argumento vai o nome da slice e depois o que vai add

	listaDeCompras = append(listaDeCompras, "banana")
	fmt.Println(listaDeCompras)

	//para add outra slice a uma slice, sem precisa chamar o índice, coloca ... após a slice
	listaDeCompras = append(listaDeCompras, listaDaIrma...)
	fmt.Println(listaDeCompras)

	var sliceNumero = []int{2, 2, 2, 2, 2, 2}
	var sliceNumero2 = []int{1, 2, 3, 4, 5, 6}
	var sliceSomaDasDuas = append(sliceNumero, sliceNumero2...)

	fmt.Println(sliceNumero)
	fmt.Println(sliceNumero2)
	fmt.Println(sliceSomaDasDuas)

	slice := make([]string, 5)
	slice[0] = "primeiro"
	fmt.Println(slice)
	tamanho := len(slice)
	capacidade := cap(slice)
	fmt.Println(tamanho, capacidade)

	//crie uma slice de tamanho 12 usando a funcao make() e atribua os meses do zodiaco indivdualmente
	//printe na tela essa slice mosntrando todos os elementos
	//printe uma fatia dessa slice do indice 2 ao oito

	meses := make([]string, 12)
	meses[0] = "janeiro"
	meses[1] = "fevereiro"
	meses[2] = "marco"
	meses[3] = "abril"
	meses[4] = "maio"
	meses[5] = "junho"
	meses[6] = "julho"
	meses[7] = "agosto"
	meses[8] = "setembro"
	meses[9] = "outubro"
	meses[10] = "novembro"
	meses[11] = "dezembro"

	fmt.Println(meses)
	fmt.Println(meses[2:8])

}
