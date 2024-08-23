package conta

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
	c.Saldo -= valor
	fmt.Printf("O valor do saque foi R$%.2f\n", valor)
	fmt.Printf("Seu saldo atual é: R$%.2f\n", c.Saldo)
	return c.Saldo
}
