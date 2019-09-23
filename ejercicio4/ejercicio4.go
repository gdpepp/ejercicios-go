package main

import "fmt"

type Asistente interface {
	GetPayment(int) int
}

type General struct {
	name string
}
type Jubilado struct {
	name string
}
type Gratis struct {
	name string
}

func (g *General) GetPayment(precio int) int {
	return precio
}
func (j Jubilado) GetPayment(precio int) int {
	return precio/2
}
func (gr Gratis) GetPayment(precio int) int {
	return 0
}

func main() {
	var total int ; precioEntrada := 10
	var asistentes[] Asistente

	asistentes = append(asistentes, new(General))
	asistentes = append(asistentes, new(General))
	asistentes = append(asistentes, new(Gratis))
	asistentes = append(asistentes, new(Jubilado))

	for _, asistente := range asistentes {
		total += asistente.GetPayment(precioEntrada)
	}

	fmt.Println("precio total: ", total)
}

