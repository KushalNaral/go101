package main

import "fmt"

type error interface {
	Error() string
}

type Nothing interface{}

type Aboutable interface {
	About() string
}

type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type MInt int

func (MInt) About() string {
	return "custom int val"
}

type Filter interface {
	About() string
	Process([]int) []int
}

type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "rem duplic numbers"
}

func (UniqueFilter) Process(inputs []int) []int {

	outs := make([]int, 0, len(inputs))
	pushed := make(map[int]bool)

	for _, n := range inputs {

		if !pushed[n] {
			pushed[n] = true
			outs = append(outs, n)
		}

	}
	return outs
}

type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiple of %v", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int {

	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

func filterAndPrint(flt Filter, unfiltered []int) []int {

	filtered := flt.Process(unfiltered)
	fmt.Println(flt.About()+":\n\t", filtered)
	return filtered
}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func mainInterfaces() {

	var a Aboutable = &Book{"Go"}
	fmt.Println(a)

	var i interface{} = &Book{"Rust"}
	fmt.Println(i)

	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}

	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	for _, ftr := range filters {
		numbers = filterAndPrint(ftr, numbers)
	}

	//reflections
	// var x interface{} = 123
	// n, ok := x.(int)
	// fmt.Println(n, ok)

	var x interface{} = DummyWriter{}
	var y interface{} = "abc"

	var w Writer
	var ok bool

	w, ok = x.(Writer)
	fmt.Println(w, ok)

	x, ok = w.(interface{})
	fmt.Println(x, ok)

	w, ok = y.(Writer)
	fmt.Println(w, ok)
	//w = y.(Writer) panics

	// switch state; v := x.(type) {
	// case tA:
	// case tB, tC:
	// case nil:
	// default:
	// }

	values := []interface{}{
		456, "abc", true, 0.33, int32(789), []int{1, 2, 3}, map[string]int{}, nil,
	}

	for _, x := range values {

		switch v := x.(type) {
		case []int:
			fmt.Println("int slice", v)
		case string:
			fmt.Println("string", v)
		case int32, float32, int:
			fmt.Println("number", v)
		case nil:
			fmt.Println("nil", v)
		default:
			fmt.Println("else", v)
		}
	}
}
