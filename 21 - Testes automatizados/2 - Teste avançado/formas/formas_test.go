package formas

import (
	"math"
	"testing"
)

type cenarioDeTesteRetangulo struct {
	base         float64
	altura       float64
	areaEsperada float64
}

type cenarioDeTesteCirculo struct {
	raio         float64
	areaEsperada float64
}

type cenarioDeTesteTriangulo struct {
	base         float64
	altura       float64
	areaEsperada float64
}

type cenarioDeTesteQuadrado struct {
	lado         float64
	areaEsperada float64
}

func truncar(numero float64) float64 {
	return math.Trunc(numero)
}
func TestArea(t *testing.T) {

	t.Run("Retângulo", func(t *testing.T) {

		cenariosDeTesteRetangulo := []cenarioDeTesteRetangulo{
			{10, 12, 120},
			{4, 5, 20},
			{7, 5, 35},
		}

		for _, cenario := range cenariosDeTesteRetangulo {
			areaEsperada := cenario.areaEsperada
			areaRecebida := cenario.altura * cenario.base
			if areaEsperada != areaRecebida {
				t.Errorf("A area do retangulo esta errada, eu esperava %f e recebi %f", areaEsperada, areaRecebida)
			}
		}
	})

	t.Run("Circulo", func(t *testing.T) {

		cenariosDeTesteCirculo := []cenarioDeTesteCirculo{
			{2, truncar(12)},
			{3, truncar(28)},
			{4, truncar(50)},
		}

		for _, cenario := range cenariosDeTesteCirculo {
			c := Circulo{Raio: cenario.raio}
			areaRecebida := truncar(c.Area())
			if cenario.areaEsperada != areaRecebida {
				t.Errorf("Area do circulo esta errada, eu esperava %f e recebi %f", cenario.areaEsperada, areaRecebida)
			}

		}
	})

	t.Run("Triangulo", func(t *testing.T) {

		cenariosDeTesteTriangulo := []cenarioDeTesteTriangulo{
			{2, 3, 3},
			{2, 4, 4},
			{3, 4, 6},
		}

		for _, cenario := range cenariosDeTesteTriangulo {
			base := cenario.base
			altura := cenario.altura
			areaEsperada := cenario.areaEsperada
			tri := Triangulo{base, altura}
			areaRecebida := tri.Area()
			if areaEsperada != areaRecebida {
				t.Errorf("Area do triangulo esta errada, eu esperava %f e recebi %f", areaEsperada, areaRecebida)
			}
		}
	})
	
	t.Run("Quadrado", func(t *testing.T) {

		cenariosDeTesteQuadrado := []cenarioDeTesteQuadrado{
			{2, 4},
			{3, 9},
			{4, 16},
			{5, 25},
		}

		for _, cenario := range cenariosDeTesteQuadrado {
			quad := Quadrado{cenario.lado}
			areaEsperada := cenario.areaEsperada
			areaRecebida := quad.Area()
			if areaEsperada != areaRecebida {
				t.Errorf("Area do quadrado esta errada, eu esperava %f e recebi %f", areaEsperada, areaRecebida)
			}
		}
	})
}
