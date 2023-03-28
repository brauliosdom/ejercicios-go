package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(valores ...int) float32 {
	minimo := 1000000
	for _, valor := range valores {
		if valor < minimo {
			minimo = valor
		}
	}
	return float32(minimo)
}

func averageFunc(valores ...int) float32 {
	var promedio int
	for i := 0; i < len(valores); i++ {
		promedio += valores[i]
	}
	return float32(promedio) / float32(len(valores))
}

func maxFunc(valores ...int) float32 {
	maximo := 0
	for _, valor := range valores {
		if valor > maximo {
			maximo = valor
		}
	}
	return float32(maximo)
}

func operation(op string) (func(valores ...int) float32, error) {
	switch op {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return nil, errors.New("Opción inválida")
	}
}

func main() {
	minFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	fmt.Println(minFunc(1, 2, 3))
	fmt.Println(averageFunc(1, 2, 3))
	fmt.Println(maxFunc(1, 2, 3))
}
