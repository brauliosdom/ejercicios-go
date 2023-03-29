package main

import "fmt"

type Product interface {
	GetPrice() float64
}

type SmallProduct struct {
	Price float64
}

type MediumProduct struct {
	Price float64
}

type LargeProduct struct {
	Price float64
}

func (s SmallProduct) GetPrice() float64 {
	return s.Price
}

func (m MediumProduct) GetPrice() float64 {
	return m.Price * 1.03
}

func (l LargeProduct) GetPrice() float64 {
	return l.Price*1.06 + 2500
}

func createProduct(typeP string, priceP float64) Product {
	switch typeP {
	case "s":
		return SmallProduct{priceP}
	case "m":
		return MediumProduct{priceP}
	case "l":
		return LargeProduct{priceP}
	default:
		return nil
	}
}

func main() {
	p1 := createProduct("s", 10.0)
	p2 := createProduct("m", 10.0)
	p3 := createProduct("l", 10.0)
	fmt.Println(p1.GetPrice(), p2.GetPrice(), p3.GetPrice())
}
