package main

import "fmt"

func main() {
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var contador int

	for _, edad := range empleados {
		if edad > 21 {
			contador++
		}
	}

	fmt.Println("Empleados mayores a 21 años:", contador)
	empleados["Federico"] = 25
	delete(empleados, "Pedro")
	fmt.Println(empleados)
}
