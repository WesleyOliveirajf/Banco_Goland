package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDaSilva := ContaCorrente{}
	contaDaSilva.titular = "Wesley Oliveira"
	contaDaSilva.saldo = 500

	fmt.Println(contaDaSilva.saldo)

	//valorSaque := 200
	//contaDaSilva.saldo = contaDaSilva - valorSaque

}
