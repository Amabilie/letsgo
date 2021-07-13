package main

import "fmt"

func main() {
	var estaEnsolarado bool
	estaEnsolarado = true

	var estaChovendo bool = false
	var estaFrio = true

	fmt.Printf("o tipo é %T, o valor é %v\n", estaEnsolarado, estaEnsolarado)
	fmt.Printf("o tipo é %T, o valor é %v\n", estaChovendo, estaChovendo)
	fmt.Printf("o tipo é %T, o valor é %v\n", estaFrio, estaFrio)

	comparacao := 10 > 2
	fmt.Println(comparacao)

	estaQuente := true
	estouDisposta := true
	estouSaindoDeCasa := false

	vouAPraia := estaQuente && estouDisposta && estouSaindoDeCasa

	fmt.Println(vouAPraia)
	fmt.Printf("%T", vouAPraia)

	//declare cinco variaveis diferentes e atribua operacoes relacionais a elas
	// use uma linha por variavel e demonstre o tipo e valor a cada uma delas

	gostoDePizza := true
	estouDeDieta := true
	estouComFome := true
	possoComerPizza := false
	temComidaEmCasa := false
	fmt.Printf("%T %V\n", gostoDePizza, gostoDePizza)
	fmt.Printf("%T %V\n", estouDeDieta, estouDeDieta)
	fmt.Printf("%T %V\n", estouComFome, estouComFome)
	fmt.Printf("%T %V\n", possoComerPizza, possoComerPizza)
	fmt.Printf("%T %V\n", temComidaEmCasa, temComidaEmCasa)

	a := 15 >= 15
	b := 100 < 1000
	c := 10 == 18
	d := 5 != 5
	e := 89 > 0 && 50 > 0

	fmt.Printf("o tipo de a é %T, e seu valor é %t\n", a, a)
	fmt.Printf("o tipo de b é %T, e seu valor é %t\n", b, b)
	fmt.Printf("o tipo de c é %T, e seu valor é %t\n", c, c)
	fmt.Printf("o tipo de d é %T, e seu valor é %t\n", d, d)
	fmt.Printf("o tipo de e é %T, e seu valor é %t\n", e, e)

}
