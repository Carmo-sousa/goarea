package goarea

import "math"

// Pi é uma constante numérica
const Pi = 3.1416

// Circ calcula a área do circulo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
