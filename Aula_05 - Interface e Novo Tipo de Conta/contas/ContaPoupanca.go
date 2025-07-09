package contas

import (
	"fmt"
	"go-orientacao-a-objetos/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (conta *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque <= conta.saldo
	if !podeSacar {
		return "saldo insuficiente"
	}
	conta.saldo -= valorSaque
	return fmt.Sprintf("Saque de R$%.2f realizado com sucesso! saldo atual: R$%.2f", valorSaque, conta.saldo)
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito >= 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Valor de depósito menor do que zero", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorTransferencia <= c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo	
}