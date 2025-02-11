package example1

type FileInvoiceDao struct {
	invoice Invoice
}

func GetFileInvoiceDao(invoice Invoice) FileInvoiceDao {
	return FileInvoiceDao{invoice: invoice}
}

func (fileInvoiceDao FileInvoiceDao) SaveInvoice() {
	// Logic to Same Invoce details here
}
