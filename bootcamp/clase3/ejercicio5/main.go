package main

import (
	"errors"
	"fmt"
)

const (
	dog = "dog"
	cat = "cat"
)

func animalDog(cantidad float64) float64 {
	return cantidad * 10
}

func animalCat(cantidad float64) float64 {
	return cantidad * 5
}

func animal(op string) (func(cantidad float64) float64, error) {
	switch op {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	default:
		return nil, errors.New("Opción inválida")
	}
}

func main() {
	animalDog, _ := animal(dog)
	animalCat, _ := animal(cat)

	var amount float64
	amount += animalDog(10)
	amount += animalCat(10)
	fmt.Println(amount)
}
