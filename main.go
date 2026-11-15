package main

import (
	"Banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	//clienteWesley := clientes.Titular{"Wesley", "123.132.123.21", "Analista"}
	//contaDoWesley := contas.ContaCorrente{clienteWesley, 123, 12456, 100}
	//fmt.Println(contaDoWesley)
	//fmt.Println(clienteWesley)

	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(200)
	PagarBoleto(&contaDenis, 60)

	fmt.Println("Seu saldo e : R$", contaDenis.ObterSaldo())

	contaDaKely := contas.ContaCorrente{}
	contaDaKely.Depositar(500)
	PagarBoleto(&contaDaKely, 300)

	fmt.Println("Seu saldo e : R$", contaDaKely.ObterSaldo())
}
