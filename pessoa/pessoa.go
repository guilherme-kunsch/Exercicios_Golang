package pessoa

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Genero string
}

// ExibirInformacoes retorna os dados da pessoa
func (c Pessoa) ExibirInformacoes() {
	fmt.Printf("Nome: %s\n", c.Nome)
	fmt.Printf("Idade: %d\n", c.Idade)
	fmt.Printf("Gênero: %s\n", c.Genero)
}

func (c Pessoa) MaiorDeIdade() {
	if c.Idade >= 18 {
		fmt.Println("A pessoa é maior de idade")
	} else {
		fmt.Println("Menor de idade")
	}
}
