package main

import (
	"fmt"
	"pratica/conta_bancaria"
)

// "pratica/conta"

func main() {
	//EXERCICIO 1
	// c := circulo.Circulo{Raio: 5}
	// circulo.EscreverArea(c)

	//EXERCICIO 2
	// minhaConta := conta.ContaBancaria{
	// 	NumeroConta: 123,
	// 	NomeTitular: "Guilherme Kunsch",
	// 	Saldo:       0.00,
	// }
	// minhaConta.Depositar(1000)
	// minhaConta.Sacar(200)

	// fmt.Println(minhaConta)

	////EXERCICIO 3
	// minhaPessoa := pessoa.Pessoa{
	// 	Nome:   "Guilherme",
	// 	Idade:  25,
	// 	Genero: "Masculino",
	// }

	// minhaPessoa.ExibirInformacoes()
	// minhaPessoa.MaiorDeIdade()

	//EXERCICIO 4
	// r := retangulo.Retangulo{Altura: 10.0, Largura: 15.0}
	// area := r.Area()
	// perimetro := r.Perimetro()

	// fmt.Printf("Área: %.2f\n", area)
	// fmt.Printf("Perímetro: %.2f\n", perimetro)

	// // Chamando o método para exibir as dimensões
	// r.ExibirDimensoes()

	////EXERCICIO 5
	// livro1 := livro.Livro{
	// 	Titulo:        "O Egito",
	// 	Autor:         "Teste Testando",
	// 	AnoPublicacao: 2024,
	// 	NumeroPaginas: 200,
	// }

	// fmt.Println("O livro está disponivel?", livro1.Disponivel())

	// fmt.Println("O livro está disponível?", livro1.Disponivel())

	// livro1.Emprestar()

	// fmt.Println("O livro está disponível?", livro1.Disponivel())

	// livro1.Devolver()

	//EXERCICIO 6
	// lamp := lampada.Lampada{
	// 	Voltagem: 220,
	// 	Potencia: 60,
	// 	Estado:   false,
	// }

	// lamp.AlterarVoltPoten(50, 110)

	// lamp.ExibirEstado()

	// lamp.Ligado()

	// lamp.ExibirEstado()

	//EXERCÍCIO 7
	// Cria novas contas com número e nome do cliente
	conta1 := conta_bancaria.NovoConta(1, "Alice")
	conta2 := conta_bancaria.NovoConta(2, "Bob")
	conta3 := conta_bancaria.NovoConta(3, "Charlie")

	// Realiza operações nas contas
	conta1.Depositar(5000.00)
	conta1.Sacar(200.00)

	conta2.Depositar(6000.00)
	conta2.Sacar(700.00)

	conta3.Depositar(700.00)
	conta3.Sacar(7000.00)

	// Exibe o relatório das contas
	fmt.Println("\nRelatório de Contas Bancárias:")
	fmt.Printf("Conta: %d | Titular: %s | Saldo Atual: R$%.2f\n", conta1.ObterNumeroConta(), conta1.ObterNomeCliente(), conta1.ExibirSaldo())
	fmt.Printf("Conta: %d | Titular: %s | Saldo Atual: R$%.2f\n", conta2.ObterNumeroConta(), conta2.ObterNomeCliente(), conta2.ExibirSaldo())
	fmt.Printf("Conta: %d | Titular: %s | Saldo Atual: R$%.2f\n", conta3.ObterNumeroConta(), conta3.ObterNomeCliente(), conta3.ExibirSaldo())

	// Imprime os extratos de cada conta
	fmt.Println("\nExtratos:")
	conta1.ImprimirExtratos()
	conta2.ImprimirExtratos()
	conta3.ImprimirExtratos()

}
