package example1

type InvoiceDao struct {
	invoice Invoice
}

func GetInvoiceDao(invoice Invoice) InvoiceDao {
	return InvoiceDao{invoice: invoice}
}

func (InvoiceDao InvoiceDao) SaveInvoice() {

}
