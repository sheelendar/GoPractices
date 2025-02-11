package example1

type Marker struct {
	name  string
	color string
	year  int
	price float32
}

func GetMarker(name, color string, year int, price float32) Marker {
	return Marker{name: name, color: color, year: year, price: price}
}
