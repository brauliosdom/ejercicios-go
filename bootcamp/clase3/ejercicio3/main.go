package main

import "fmt"

func calculaSalario(minutos int, categoria string) float32 {
	var salario float32
	horas := minutos / 60
	switch categoria {
	case "A":
		salario = float32(horas*3000) * 1.5
	case "B":
		salario = float32(horas*1500) * 1.2
	case "C":
		salario = float32(horas * 1000)
	}
	return salario
}

func main() {
	fmt.Println(calculaSalario(120, "C"))
}
