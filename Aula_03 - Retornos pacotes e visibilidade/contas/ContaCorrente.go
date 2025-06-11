package contas

import (
	"fmt"
)

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (conta *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= conta.Saldo
	if !podeSacar {
		return "Saldo insuficiente"
	}
	conta.Saldo -= valorSaque
	return fmt.Sprintf("Saque de R$%.2f realizado com sucesso! Saldo atual: R$%.2f", valorSaque, conta.Saldo)
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito >= 0 {
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.Saldo
	} else {
		return "Valor de depósito menor do que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia <= c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}