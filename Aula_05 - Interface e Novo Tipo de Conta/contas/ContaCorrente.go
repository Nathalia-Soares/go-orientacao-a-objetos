package contas

import (
	"fmt"
	"go-orientacao-a-objetos/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= conta.saldo
	if !podeSacar {
		return "saldo insuficiente"
	}
	conta.saldo -= valorSaque
	return fmt.Sprintf("Saque de R$%.2f realizado com sucesso! saldo atual: R$%.2f", valorSaque, conta.saldo)
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito >= 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Valor de depósito menor do que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia <= c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo	
}