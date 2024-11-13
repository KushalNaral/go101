package main

import "fmt"

func MainStruct() {

	type Book struct {
		Pages int
	}

	p := &Book{100}
	p.Pages = 200

	//here why we dont ened *p.Pages is cause it goes to become
	// (*p).Pages which gets auto dereferenced by the compiler

	fmt.Println(p)

	type C int

	type A struct {
		b *C
	}

}
