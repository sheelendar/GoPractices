package main

import "fmt"

/*
Requirements
▪ All books should have attributes: Name, Author, Publisher,
Publish Year, Category, Price, Count (sold)
The following should be implemented on top of this catalog
▪ Add book to the catalog
▪ Search a book by partial book/author name
▪ Get most sold books, by author name/category, limit
Maintain the system on memory (no files/DB
*/

func main() {
	book1 := GetBookConstractor("HP & The PS", "J K Rowling", "Bloomsbury", 1997, 80, 200.0, FICTION)
	book2 := GetBookConstractor("HP & The COS", "J K Rowling", "Bloomsbury", 1998, 100, 1000.0, FICTION)

	book3 := GetBookConstractor("HP & The POA", "J K Rowling", "Bloomsbury", 1999, 500, 2000.0, FICTION)

	book4 := GetBookConstractor("HP & The HBP", "J K Rowling", "Bloomsbury", 2005, 700, 3000.0, FICTION)

	book5 := GetBookConstractor("The Immortals of Meluha", "Mohan", "Westland", 2010, 600, 1500.0, MYSTERY)

	book6 := GetBookConstractor("The Secret of the Nagas", "Mohan", "Westland", 2011, 400, 2500.0, MYSTERY)

	book7 := GetBookConstractor("The Oath of the Vayuputras", "Mohan", "Westland", 2013, 200, 3500.0, MYSTERY)

	book8 := GetBookConstractor("Do Androids drems of Electric Sheep", "philips k", "BoubleDay", 1968, 20, 30.0, SCI_FI)

	catalog := GetCatalogConstractor()
	catalog.addBooktoCatalog(book1)
	catalog.addBooktoCatalog(book2)
	catalog.addBooktoCatalog(book3)
	catalog.addBooktoCatalog(book4)
	catalog.addBooktoCatalog(book5)
	catalog.addBooktoCatalog(book6)
	catalog.addBooktoCatalog(book7)
	catalog.addBooktoCatalog(book8)

	fmt.Println("most book sold of auther ", book1.author)
	books := catalog.MostSoldBooksByWriter(book1.author, 3)
	displayBooks(books)
	fmt.Println("***********************************************")

	catalog.displayBookByAuthor(book1.author)
	fmt.Println("***********************************************")

	catalog.displayBookByCategory(MYSTERY)
}
