package formas

import (
	"math"
)

// Interface que define que uma forma sempre vai ter o metodo Area
type Forma interface {
	Area() float64
}

// Metodo para cada forma
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pow(c.Raio, 2) * math.Pi
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (q Quadrado) Area() float64 {
	return math.Pow(q.Lado, 2)
}

// struct que define a necessidade de cada forma
type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Quadrado struct {
	Lado float64
}

// Função que escreve a Area, e recebe como parametro a forma

// r := retangulo{10, 15}
// c := circulo{10}
// q := quadrado{2}
// t := triangulo{2, 3}
// escreverArea(r)
// escreverArea(c)
// escreverArea(q)
// escreverArea(t)
