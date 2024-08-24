package main

import (
	"pratica/lampada"
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
	lamp := lampada.Lampada{
		Voltagem: 220,
		Potencia: 60,
		Estado:   false,
	}

	lamp.AlterarVoltPoten(50, 110)

	lamp.ExibirEstado()

	lamp.Ligado()

	lamp.ExibirEstado()
}
