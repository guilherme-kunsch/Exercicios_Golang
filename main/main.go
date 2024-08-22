package main

import (
	// "pratica/conta"
	"pratica/pessoa"
)

func main() {
	// c := circulo.Circulo{Raio: 5}
	// circulo.EscreverArea(c)

	// minhaConta := conta.ContaBancaria{
	// 	NumeroConta: 123,
	// 	NomeTitular: "Guilherme Kunsch",
	// 	Saldo:       0.00,
	// }
	// minhaConta.Depositar(1000)
	// minhaConta.Sacar(200)

	// fmt.Println(minhaConta)

	minhaPessoa := pessoa.Pessoa{
		Nome:   "Guilherme",
		Idade:  25,
		Genero: "Masculino",
	}

	minhaPessoa.ExibirInformacoes()
	minhaPessoa.MaiorDeIdade()
}
