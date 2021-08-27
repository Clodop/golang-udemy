package goarea

import "math"

//é obrigatorio coentar variaveis publicas
const Pi = 3.1416

//calcaula circuferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula ret....
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
