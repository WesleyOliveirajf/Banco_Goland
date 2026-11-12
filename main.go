package main

import "fmt"

func main() {
	contaDaSilva := ContaCorrente{titular: "Silva", saldo: 300}
	contaDoWesley := ContaCorrente{titular: "Wesley", saldo: 100}

	status := contaDaSilva.Transferir(200, &contaDoWesley)

	fmt.Println(status)
	fmt.Println(contaDaSilva)
	fmt.Println(contaDoWesley)

}
