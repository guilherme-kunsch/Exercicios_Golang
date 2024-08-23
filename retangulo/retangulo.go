package retangulo

import "fmt"

type Retangulo struct {
	Altura  float64
	Largura float64
}

func (c Retangulo) Area() float64 {
	return c.Altura * c.Largura
}

func (c Retangulo) Perimetro() float64 {
	return (c.Altura * 2) + (c.Largura * 2)
}

func (r Retangulo) ExibirDimensoes() {
	fmt.Printf("Altura: %.2f, Largura: %.2f\n", r.Altura, r.Largura)
}
