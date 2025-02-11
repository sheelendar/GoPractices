package example1

type InvoicePrinter struct {
	invoice Invoice
}

func GetInvoicePrinter(invoice Invoice) InvoicePrinter {
	return InvoicePrinter{invoice: invoice}
}

func (invoiceDao InvoicePrinter) DisplayInvoice() {
	// Logic to display Invoce details here
}
