package interfaces

import (
	"fmt"
	"math"
)

type forma interface {
	area () float64
}

func EscreverArea (f forma) {
	fmt.Printf("A area da forma é %0.2f", f.area())
}

type Retangulo struct{
	Altura float64
	Largura float64
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) area() float64{
	return math.Pi * (c.Raio * c.Raio)
}


