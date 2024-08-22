package main

import (
	// "pratica/circulo"
	"fmt"
	"pratica/conta"
)

func main() {
	// c := circulo.Circulo{Raio: 5}
	// circulo.EscreverArea(c)

	minhaConta := conta.ContaBancaria{
		NumeroConta: 123,
		NomeTitular: "Guilherme Kunsch",
		Saldo:       0.00,
	}
	minhaConta.Depositar(1000)
	minhaConta.Sacar(200)

	fmt.Println(minhaConta)
}
