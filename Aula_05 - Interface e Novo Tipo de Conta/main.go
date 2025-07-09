package main

import (
	"fmt"
	"go-orientacao-a-objetos/clientes"
	"go-orientacao-a-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valorSaque float64) string
}

func main() {
	titular := clientes.Titular{
		Nome:      "Eni Essi",
		Cpf:       "123.456.789-00",
		Profissao: "Arquiteta",
	}

	contaCorrente := contas.ContaCorrente{
		Titular:       titular,
		NumeroAgencia: 1234,
		NumeroConta:   56789,
	}

	contaPoupanca := contas.ContaPoupanca{
		Titular:       titular,
		NumeroAgencia: 1234,
		NumeroConta:   56789,
		Operacao:      001,
	}

	contaCorrente.Depositar(1000.00)
	contaPoupanca.Depositar(100.00)

	PagarBoleto(&contaCorrente, 200.00)
	PagarBoleto(&contaPoupanca, 50.00)

	fmt.Println(contaCorrente.ObterSaldo())
	fmt.Println(contaPoupanca.ObterSaldo())
}
