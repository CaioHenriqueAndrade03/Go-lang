package main


// Interface que define que uma forma sempre vai ter o metodo area
type forma interface {
	area() float64
}

// Metodo para cada forma
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c circulo) area() float64 {
	return math.Pow(c.raio, 2) * math.Pi
}

func (t triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

func (q quadrado) area() float64 {
	return math.Pow(q.lado, 2)
}

// struct que define a necessidade de cada forma
type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

type triangulo struct {
	base   float64
	altura float64
}

type quadrado struct {
	lado float64
}

// Função que escreve a area, e recebe como parametro a forma
func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())

}

func main() {
	r := retangulo{10, 15}
	c := circulo{10}
	q := quadrado{2}
	t := triangulo{2, 3}
	escreverArea(r)
	escreverArea(c)
	escreverArea(q)
	escreverArea(t)

}

}