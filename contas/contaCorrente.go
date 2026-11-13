package contas

import "Banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Transferir(valordaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valordaTransferencia <= c.saldo && valordaTransferencia > 0 {
		c.saldo -= valordaTransferencia
		contaDestino.Depositar(valordaTransferencia)
		return true

	} else {
		return false
	}

}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insufuciente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depositor Realizado com Suscesso", c.saldo
	} else {
		return "deposito Nao Efetuado", c.saldo
	}

}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
