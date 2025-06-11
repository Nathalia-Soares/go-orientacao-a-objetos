package main

import (
	"go-orientacao-a-objetos/contas"
	"fmt"
)

func main() {

	contaFulano := contas.ContaCorrente{
		Titular: "Fulano",
		Saldo:   300,
	}

	contaSiclano := contas.ContaCorrente{
		Titular: "Siclano",
		Saldo:   100,
	}

	statusTransferencia := contaFulano.Transferir(50, &contaSiclano)

	var contaDaNathalia *contas.ContaCorrente
	contaDaNathalia = new(contas.ContaCorrente)
	contaDaNathalia.Titular = "Nathalia"
	contaDaNathalia.Saldo = 500

	fmt.Println(contaDaNathalia.Saldo)
	fmt.Println(contaDaNathalia.Sacar(100.0))
	status, valor := contaDaNathalia.Depositar(500.0)
	fmt.Println(status, valor)
	fmt.Println("-------------------------------------------")
	fmt.Println(statusTransferencia)
	fmt.Println(contaFulano)
	fmt.Println(contaSiclano)
}
