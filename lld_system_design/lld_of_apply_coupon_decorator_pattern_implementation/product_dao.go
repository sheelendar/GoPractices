package main

type ProductType int

const (
	Electronic = iota + 1
	Funiture
	Decorative
)

// **************************************************** ProductPrice interface *******************************
type ProductPrice interface {
	getPrice() float64
}

// **************************************************** Product *******************************
type Product struct {
	name          string
	originalPrice float64
	productType   ProductType
}

// **************************************************** Item1 *******************************
type Item1 struct {
	Product
}

func Item1Constractor(name string, originalPrice float64, productType ProductType) *Item1 {
	return &Item1{Product: Product{name: name, originalPrice: originalPrice, productType: productType}}
}
func (item *Item1) getPrice() float64 {
	return item.originalPrice
}

// **************************************************** Item2 *******************************
type Item2 struct {
	Product
}

func Item2Constractor(name string, originalPrice float64, productType ProductType) *Item2 {
	return &Item2{Product: Product{name: name, originalPrice: originalPrice, productType: productType}}
}
func (item *Item2) getPrice() float64 {
	return item.originalPrice
}
