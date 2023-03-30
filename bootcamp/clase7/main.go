package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type DatoVacioError struct {
	msg string
}

func (e *DatoVacioError) Error() string {
	return fmt.Sprintf("error: el dato '%s' esta vacio", e.msg)
}

type Cliente struct {
	Legajo    string
	Nombre    string
	DNI       string
	Telefono  string
	Domicilio string
}

// TODO: obtener registros de un archivo
var Clientes = []Cliente{
	{"legajo", "nombre", "dni", "telefono", "domicilio"},
}

func (c Cliente) VerificaCliente() error {
	for _, cliente := range Clientes {
		if c.DNI == cliente.DNI {
			return errors.New("error: el cliente ya existe")
		}
	}
	return nil
}

func (c Cliente) VerificaDatos() error {
	switch {
	case c.Legajo == "":
		return &DatoVacioError{"legajo"}
	case c.Nombre == "":
		return &DatoVacioError{"nombre"}
	case c.DNI == "":
		return &DatoVacioError{"dni"}
	case c.Telefono == "":
		return &DatoVacioError{"telefono"}
	case c.Domicilio == "":
		return &DatoVacioError{"domicilio"}
	default:
		return nil
	}
}

// TODO: guardar registro en el archivo
func (c Cliente) AgregaCliente() error {
	Clientes = append(Clientes, c)
	return nil
}

func main() {
	f, err := os.Open("customers.txt")

	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}

	text, err := io.ReadAll(f)
	fmt.Println(string(text))

	if err != nil {
		panic("error al intentar leer el archivo")
	}

	c := Cliente{"legajo", "nombre", "dni", "telefono", ""}

	// 1. verifica cliente unico
	err = c.VerificaCliente()
	if err != nil {
		panic(err)
	}

	// 2. verifica datos completos
	err = c.VerificaDatos()
	if err != nil {
		panic(err)
	}

	// 3. guarda info del cliente
	err = c.AgregaCliente()
	if err != nil {
		panic(err)
	}

	defer func() {
		// TODO: recuperar todos los panics
		er := recover()
		if er != nil {
			panic(er)
		}
		fmt.Println("ejecución finalizada")
		f.Close()
	}()

}
