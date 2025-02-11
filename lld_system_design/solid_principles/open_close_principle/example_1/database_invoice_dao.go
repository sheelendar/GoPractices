package example1

type DatabaseInvoiceDao struct {
	invoice Invoice
}

func GetDataBaseInvoiceDao(invoice Invoice) DatabaseInvoiceDao {
	return DatabaseInvoiceDao{invoice: invoice}
}

func (databaseInvoiceDao DatabaseInvoiceDao) SaveInvoice() {
	// Logic to Save Invoce details here
}
