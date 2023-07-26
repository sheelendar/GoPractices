package main

import (
	"log"
	"sort"
)

type Planet struct {
	Name       string  `json:"name"`
	Aphelion   float64 `json:"aphelion"`   // in million km
	Perihelion float64 `json:"perihelion"` // in million km
	Axis       int64   `json:"Axis"`       // in km
	Radius     float64 `json:"radius"`
}

func main() {
	var mars = new(Planet)
	mars.Name = "Mars"
	mars.Aphelion = 249.2
	mars.Perihelion = 206.7
	mars.Axis = 227939100
	mars.Radius = 3389.5

	var earth = new(Planet)
	earth.Name = "Earth"
	earth.Aphelion = 151.930
	earth.Perihelion = 147.095
	earth.Axis = 149598261
	earth.Radius = 6371.0

	var venus = new(Planet)
	venus.Name = "Venus"
	venus.Aphelion = 108.939
	venus.Perihelion = 107.477
	venus.Axis = 108208000
	venus.Radius = 6051.8

	planets := []Planet{*mars, *venus, *earth}
	sort.Slice(planets, func(i, j int) bool {
		return planets[i].Axis > planets[j].Axis
	})
	log.Println(planets)
}
