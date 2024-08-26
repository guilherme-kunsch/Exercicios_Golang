package lampada

import (
	"fmt"
	"math/rand"
	"time"
)

type Lampada struct {
	Voltagem float64
	Potencia float64
	Estado   bool
	Queimada bool
}

func (l *Lampada) AlterarVoltPoten(valorPotencia, valorVoltagem float64) {
	l.Voltagem = valorVoltagem
	l.Potencia = valorPotencia
	fmt.Printf("O valor de potência: %.2f\n", l.Potencia)
	fmt.Printf("O valor de voltagem: %.2f\n", l.Voltagem)
}

func (l *Lampada) LerVoltagem() float64 {
	return l.Voltagem
}

func (l *Lampada) LerPotencia() float64 {
	return l.Potencia
}

func (l *Lampada) Ligado() bool {
	if rand.Float64() <= 0.15 {
		l.Queimada = true
		fmt.Println("A lâmpada queimou ao ser ligada!")
		l.Estado = false
		return false
	}

	l.Estado = true
	return l.Estado
}

func (l *Lampada) Desligado() bool {
	l.Estado = false
	return l.Estado
}

func (l *Lampada) ExibirEstado() {
	estado := "desligada"
	if l.Estado {
		estado = "ligada"
	}
	if l.Queimada {
		estado = "queimada"
	}
	fmt.Printf("Lâmpada está %s, Potência: %.2fW, Voltagem: %.2fV\n", estado, l.Potencia, l.Voltagem)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//a lampada tem 15% de chances de ser queimada ao ser ligada
