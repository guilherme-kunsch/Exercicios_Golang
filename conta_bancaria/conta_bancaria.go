// conta_bancaria/conta_bancaria.go
package conta_bancaria

import (
	"fmt"
	"time"
)

// ContaBancaria representa uma conta com saldo, extratos, número da conta e nome do cliente.
type ContaBancaria struct {
	Saldo       float64
	Extratos    [1000]Extrato
	Indice      int
	NumeroConta int
	NomeCliente string
}

func NovoConta(numero int, nome string) *ContaBancaria {
	return &ContaBancaria{
		NumeroConta: numero,
		NomeCliente: nome,
	}
}

func (c *ContaBancaria) Depositar(valor float64) {
	c.Saldo += valor
	c.adicionarExtrato(valor)
}

func (c *ContaBancaria) Sacar(valor float64) {
	if valor > c.Saldo {
		fmt.Println("Saldo insuficiente")
		return
	}
	c.Saldo -= valor
	c.adicionarExtrato(-valor)
}

func (c *ContaBancaria) adicionarExtrato(valor float64) {
	if c.Indice >= len(c.Extratos) {
		fmt.Println("O vetor de extratos está cheio")
		return
	}
	c.Extratos[c.Indice] = Extrato{
		Data:  time.Now(),
		Valor: valor,
	}
	c.Indice++
}

func (c *ContaBancaria) ImprimirExtratos() {
	fmt.Printf("Extratos da Conta: %d | Nome Cliente: %s | Saldo atual: R$%.2f\n", c.NumeroConta, c.NomeCliente, c.Saldo)
	for i := 0; i < c.Indice; i++ {
		extrato := c.Extratos[i]
		fmt.Printf("Data: %s | Valor: R$%.2f\n", extrato.Data.Format("02/01/2006 15:04:05"), extrato.Valor)
	}
}

func (c *ContaBancaria) ExibirSaldo() float64 {
	return c.Saldo
}

func (c *ContaBancaria) ObterNumeroConta() int {
	return c.NumeroConta
}

func (c *ContaBancaria) ObterNomeCliente() string {
	return c.NomeCliente
}
