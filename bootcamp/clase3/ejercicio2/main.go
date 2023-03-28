package main

import (
	"errors"
	"fmt"
)

func calculaPromedio(valores ...float32) (float32, error) {
	var promedio float32
	for i := 0; i < len(valores); i++ {
		if valores[i] < 0 {
			return 0, errors.New("Valor negativo")
		}
		promedio += valores[i]
	}
	return promedio / float32(len(valores)), nil
}

func main() {
	fmt.Println(calculaPromedio(1, 2, 3))
}
