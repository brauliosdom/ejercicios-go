package main

import "fmt"

func main() {
	edad := 20
	esEmpleado := false
	añosTrabajando := 0

	if edad < 22 {
		fmt.Println("Debe ser mayor de edad")
		return
	}
	if esEmpleado != true {
		fmt.Println("Debe ser empleado")
		return
	}
	if añosTrabajando < 1 {
		fmt.Println("Debe tener 1 año trabajando")
		return
	}

	fmt.Println("Prestamo otorgado")
}
