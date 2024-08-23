package main

// "pratica/conta"
import (
	"fmt"
	"pratica/retangulo"
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

	// minhaPessoa := pessoa.Pessoa{
	// 	Nome:   "Guilherme",
	// 	Idade:  25,
	// 	Genero: "Masculino",
	// }

	// minhaPessoa.ExibirInformacoes()
	// minhaPessoa.MaiorDeIdade()

	r := retangulo.Retangulo{Altura: 10.0, Largura: 15.0}
	area := r.Area()
	perimetro := r.Perimetro()

	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)

	// Chamando o método para exibir as dimensões
	r.ExibirDimensoes()
}
