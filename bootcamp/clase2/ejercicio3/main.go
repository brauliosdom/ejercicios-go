package main

import "fmt"

func main() {
	dia, mes := 29, "febrero"
	var meses = map[string]int{
		"enero":      31,
		"febrero":    28,
		"marzo":      31,
		"abril":      30,
		"mayo":       31,
		"junio":      30,
		"julio":      31,
		"agosto":     31,
		"septiembre": 30,
		"octubre":    31,
		"noviembre":  30,
		"diciembre":  31,
	}

	if dia > meses[mes] {
		fmt.Println("Invalido")
	} else {
		fmt.Println("Valido")
	}
}
