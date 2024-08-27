package conta_bancaria

import "fmt"

type ContaBancaria struct {
	NumeroConta int
	NomeTitular string
	Saldo       float64
}

// métodos

func (c *ContaBancaria) Depositar(valor float64) float64 {
	c.Saldo += valor
	fmt.Printf("Foi depositado o valor de R$%.2f\n", valor)
	fmt.Printf("Seu saldo atual é: R$%.2f\n", c.Saldo)
	return c.Saldo
}

func (c *ContaBancaria) Sacar(valor float64) float64 {
	if valor > c.Saldo {
		fmt.Printf("Saldo insuficiente para saque na conta %d de %s.\n", c.NumeroConta, c.NomeTitular)
	} else {
		c.Saldo -= valor
		fmt.Printf("Saque de %.2f realizado com sucesso na conta %d de %s. Seu saldo é R$%.2f\n", valor, c.NumeroConta, c.NomeTitular, c.Saldo)
	}
	return c.Saldo
}

func (c *ContaBancaria) ObterSaldo() float64 {
	return c.Saldo
}

func (c *ContaBancaria) ObterNumeroConta() int {
	return c.NumeroConta
}

func (c *ContaBancaria) ObterNomeCliente() string {
	return c.NomeTitular
}
