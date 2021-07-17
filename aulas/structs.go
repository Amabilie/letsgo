package main

import "fmt"

//struct é tipo um "objeto"

type Apartamento struct {
	numero            int
	nomeProprietaria  string
	possuiVagaGaragem bool
}

func main() {
	ap1 := Apartamento{
		numero:            1301,
		nomeProprietaria:  "Amabilie",
		possuiVagaGaragem: true,
	}
	fmt.Println(ap1)
	fmt.Printf("O apartamento n %d tem como prietaria a %s. Ele tem vaga de garagem? %t", ap1.numero, ap1.nomeProprietaria, ap1.possuiVagaGaragem)

}
