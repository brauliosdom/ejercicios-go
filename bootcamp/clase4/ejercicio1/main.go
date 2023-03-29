package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var Products = []Product{
	{1, "hola", 10.0, "mundo", "ok"},
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) Product {
	for _, p := range Products {
		if p.ID == id {
			return p
		}
	}
	return Product{}
}

func main() {
	p := Product{2, "hola", 10.0, "mundo", "ok"}
	p.Save()
	p.GetAll()
	fmt.Println(getById(1))
	fmt.Println(getById(2))
}
