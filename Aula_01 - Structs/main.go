package main

import (
    "fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	conta1 := ContaCorrente{
		titular: "Nathalia",
		numeroAgencia: 12345,
		numeroConta: 98765,
		saldo: 250.63,
	}

	conta2 := ContaCorrente{
		"Teste",
		456789,
		111111,
		300.98,
	}

	fmt.Println(conta1)
	fmt.Println(conta2)
}