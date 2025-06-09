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

func (conta *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= conta.saldo
	if !podeSacar {
		return "Saldo insuficiente"
	}
	conta.saldo -= valorSaque
	return fmt.Sprintf("Saque de R$%.2f realizado com sucesso! Saldo atual: R$%.2f", valorSaque, conta.saldo)
}

func main() {
	conta1 := ContaCorrente{
		titular:       "Nathalia",
		numeroAgencia: 12345,
		numeroConta:   98765,
		saldo:         250.63,
	}

	conta2 := ContaCorrente{
		"Teste",
		456789,
		111111,
		300.98,
	}

	fmt.Println(conta1)
	fmt.Println(conta2)

	var contaDaNathalia *ContaCorrente
	contaDaNathalia = new(ContaCorrente)
	contaDaNathalia.titular = "Nathalia"
	contaDaNathalia.saldo = 500

	fmt.Println(*contaDaNathalia)

	fmt.Println(conta1 == conta2)

	fmt.Println(contaDaNathalia.saldo)
	fmt.Println(contaDaNathalia.Sacar(100.0))
}
