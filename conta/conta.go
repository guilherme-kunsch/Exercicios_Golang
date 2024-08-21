package conta

import "fmt"

type ContaBancaria struct {
	numeroConta float64
	nomeTitular string
	saldo       float64
}

type exibirResultados interface {
	depositar() float64
}

//métodos

func (c ContaBancaria) depositar(valor float64) {
	c.saldo = c.saldo + valor
	fmt.Printf("Foi depositado o valor de %f", valor)
}
