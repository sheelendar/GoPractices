package main

type Viecal interface {
	Shape(a, b int) int
	Area(a, b int) int
}

type Car struct {
	a int
	b int
}

func (car Car) Shape(a, b int) int {
	return 0
}

func (car Car) Area(a, b int) int {
	return 0
}
