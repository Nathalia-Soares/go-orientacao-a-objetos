package main

import (
	"fmt"
	"go-orientacao-a-objetos/clientes"
	"go-orientacao-a-objetos/contas"
)

func main() {
	titular := clientes.Titular{
		Nome:      "Eni Essi",
		Cpf:       "123.456.789-00",
		Profissao: "Arquiteta",
	}

	conta := contas.ContaCorrente{
		Titular:       titular,
		NumeroAgencia: 1234,
		NumeroConta:   56789,
	}

	conta.Depositar(1000.00)

	fmt.Println(conta.ObterSaldo())
}
