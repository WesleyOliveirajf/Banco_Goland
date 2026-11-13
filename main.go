package main

import (
	"Banco/contas"
	"fmt"
)

func main() {
	//clienteWesley := clientes.Titular{"Wesley", "123.132.123.21", "Analista"}
	//contaDoWesley := contas.ContaCorrente{clienteWesley, 123, 12456, 100}
	//fmt.Println(contaDoWesley)
	//fmt.Println(clienteWesley)

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(-100)

	fmt.Println(contaExemplo.ObterSaldo())
}
