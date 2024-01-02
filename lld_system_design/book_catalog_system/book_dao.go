package main

import (
	"fmt"
	"sort"
)

type Category int

const (
	FICTION = iota + 1
	SCI_FI
	MYSTERY
	FABLE
	MYTHOLOGY
)

type Book struct {
	name        string
	author      string
	Publisher   string
	publishYear int
	category    Category
	price       float64
	count       int
}

func (book *Book) DisplayBookDetails() {
	fmt.Println("Book Author ", book.author, " Book name ", book.name, " publish Year ", book.publishYear,
		" book price ", book.price, " sold count ", book.count)
}
func GetBookConstractor(name, author, publisher string, publishYear, count int, price float64, category Category) *Book {
	return &Book{name: name, author: author, Publisher: publisher,
		publishYear: publishYear, price: price, category: category, count: count}
}

//****************************************** Catalog ******************************

type Catalog struct {
	books           []*Book
	booksByAuthor   map[string][]*Book
	booksByCaterogy map[Category][]*Book
}

func GetCatalogConstractor() *Catalog {
	return &Catalog{booksByAuthor: map[string][]*Book{}, booksByCaterogy: make(map[Category][]*Book)}
}

func (catalog *Catalog) addBooktoCatalog(book *Book) {
	catalog.books = append(catalog.books, book)
	catalog.booksByAuthor[book.author] = append(catalog.booksByAuthor[book.author], book)
	catalog.booksByCaterogy[book.category] = append(catalog.booksByAuthor[book.author], book)
}

func (catalog *Catalog) searchBooKByName(name string) []*Book {
	var res []*Book
	for _, book := range catalog.books {
		if book.name == name {
			res = append(res, book)
		}
	}
	return res
}

func (catalog *Catalog) MostSoldBooksByWriter(name string, limit int) []*Book {
	boolList := catalog.booksByAuthor[name]
	sort.Slice(boolList, func(i, j int) bool {
		return boolList[i].count > boolList[j].count
	})
	if limit < len(boolList) {
		return boolList[:limit]
	}
	return boolList
}

func (catalog *Catalog) searchBoolByAuthor(name string) []*Book {
	return catalog.booksByAuthor[name]
}

func (catalog *Catalog) displayBookByAuthor(name string) []*Book {
	boolList := catalog.booksByAuthor[name]
	for _, book := range boolList {
		book.DisplayBookDetails()
	}
	return boolList
}

func (catalog *Catalog) displayBookByCategory(category Category) []*Book {
	boolList := catalog.booksByCaterogy[category]
	for _, book := range boolList {
		book.DisplayBookDetails()
	}
	return boolList
}

func displayBooks(books []*Book) {
	for _, book := range books {
		book.DisplayBookDetails()
	}
}
