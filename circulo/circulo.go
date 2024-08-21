package circulo

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
	perimetro() float64
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func (c Circulo) perimetro() float64 {
	return (math.Pi * c.Raio) * 2
}

func EscreverArea(f forma) {
	fmt.Printf("A área do circulo é %0.2f\n", f.area())
	fmt.Printf("O perímetro do circulo é %0.2f\n", f.perimetro())

	if circulo, ok := f.(Circulo); ok {
		fmt.Printf("O raio é %0.2f\n", circulo.Raio)
	} else {
		fmt.Println("Não foi possível obter o raio.")
	}
}
