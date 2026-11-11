package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insufuciente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
	}

}

func main() {
	contaDaSilva := ContaCorrente{}
	contaDaSilva.titular = "silvia"
	contaDaSilva.saldo = 500
	fmt.Println("saldo em conta: R$", contaDaSilva.saldo)

	contaDaSilva.Depositar(-200)
	fmt.Println("saldo em conta: R$", contaDaSilva.saldo)
}
