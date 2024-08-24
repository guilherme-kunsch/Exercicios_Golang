package livro

import "fmt"

type Livro struct {
	Titulo        string
	Autor         string
	AnoPublicacao int
	NumeroPaginas int
	Livre         bool
}

func (l *Livro) Emprestar() {
	l.Livre = true
	fmt.Println("Livro emprestado")
}

func (l *Livro) Devolver() {
	l.Livre = false
	fmt.Println("Livro devolvido")
}

func (l Livro) Disponivel() bool {
	return !l.Livre
}
