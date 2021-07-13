package main

import "fmt"

func main() {

	var pesoDoTomate, quantidadeDeGarrafasAzeite, unidadesDeSabonete float64
	pesoDoTomate = 2.600
	quantidadeDeGarrafasAzeite = 2
	unidadesDeSabonete = 7

	var precoTomate, precoAzeite, precoSabonete float64
	precoTomate = 10.3
	precoAzeite = 19
	precoSabonete = 5.80

	valorTotalCompra := (pesoDoTomate * precoTomate) + (quantidadeDeGarrafasAzeite * precoAzeite) + (unidadesDeSabonete * precoSabonete)

	fmt.Printf("O valor da compra foi: \n %v, \n\t só uma garrafa de azeite já custou %v", valorTotalCompra, precoAzeite)
}
