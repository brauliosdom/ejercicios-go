package main

import "fmt"

func calculaImpuesto(sueldo float32) float32 {
	var impuesto float32
	if sueldo > 150000 {
		impuesto = sueldo * 0.27
	} else if sueldo > 50000 {
		impuesto = sueldo * 0.17
	}
	return impuesto
}

func main() {
	fmt.Println(calculaImpuesto(50001))
}
