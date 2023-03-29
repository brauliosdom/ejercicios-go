package main

import "fmt"

type AlumnoDetalle interface {
	details()
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) details() {
	fmt.Println("Nombre:", a.Nombre)
	fmt.Println("Apellido:", a.Apellido)
	fmt.Println("DNI:", a.DNI)
	fmt.Println("Fecha:", a.Fecha)
}

func main() {
	alumno := Alumno{
		Nombre:   "Braulio",
		Apellido: "Dominguez",
		DNI:      "DNI",
		Fecha:    "16-04-1999",
	}

	alumno.details()
}
