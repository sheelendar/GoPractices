package example1

type Invoice struct {
	marker   Marker
	quantity int
}

func GetInvoice(marker Marker, quantity int) Invoice {
	return Invoice{marker: marker, quantity: quantity}
}

func (invoice Invoice) CalculateTotal() float32 {
	price := invoice.marker.price * float32(invoice.quantity)
	return price
}
