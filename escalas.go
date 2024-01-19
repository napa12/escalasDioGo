package main

import "fmt"

func main() {

	var ebulicaoCelsius float64 = 100.0
	var ebulicaoKelvin float64 = ebulicaoCelsius + 273.15
	fmt.Printf("O temperatura de ebulição da água em graus Celsius é: %v\nE o valor da temperatura de ebulição da água em graus Kelvin é: %v", ebulicaoCelsius, ebulicaoKelvin)
}
