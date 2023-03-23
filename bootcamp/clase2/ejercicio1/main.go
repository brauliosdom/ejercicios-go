package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	palabra := bufio.NewScanner(os.Stdin)
	for palabra.Scan() {
		if palabra.Text() == "exit" {
			return
		}
		fmt.Println("Tamaño de la palabra:", len(palabra.Text()))
		for i, letra := range palabra.Text() {
			fmt.Printf("Letra %d: %c\n", i+1, letra)
		}
	}
}
