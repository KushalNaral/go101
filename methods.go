package main

import "fmt"

type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}

func (age *Age) Increase() {
	*age++
}

type FilterFunc func(in int) bool

func (ff FilterFunc) Filte(in int) bool {
	return ff(in)
}

type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, preset := ss[key]
	return preset
}

func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}

func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// type Book struct {
// 	pages int
// }

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

// func (b *Book) SetPages(pages int) {
// 	(*Book).SetPages(b, pages)
// }

type Books []Book

func (books *Books) Modify() {
	*books = append(*books, Book{2803})
	(*books)[0].pages = 500
}

type MyInt int

func mainMethods() {

	var book Book

	// book.SetPages(21)
	// book.SetPages(22)

	(*Book).SetPages(&book, 123)

	//func(b *Book) Pages = (*b).Pages

	fmt.Printf("\n%v\n", book)

	fmt.Printf("%T\n", book.pages)
	fmt.Printf("%T\n", (&book).SetPages)
	fmt.Printf("%T\n", (&book).Pages)

	(&book).SetPages(1231930)
	book.SetPages(1231)
	fmt.Println(book.Pages())
	fmt.Println((&book).Pages())

	fmt.Println((StringSet(nil)).Has)

	var b Book
	b.SetPages(12312312312)
	fmt.Println(b.pages)

	var books = Books{{123123123}, {21}}
	books.Modify()
	fmt.Println(books)
}
