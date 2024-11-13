package main

import "fmt"

func Sum(values ...int64) (sum int64) {

	sum = 0

	for _, v := range values {
		sum += v
	}

	return
}

func Concat(sep string, tokens ...string) (s string) {

	s = ""

	for i, b := range tokens {

		if i != 0 {
			s += sep
		}

		s += b
	}

	return

}

// func Print(a ...interface{}) (n int, err error)
//
// func Printf(format string, a ...interface{}) (n int, err error)
//
// func Println(a ...interface{}) (n int, err error)

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n)
}

func mainFunc() {

	// func(values ...int64) (sum int64)
	// func(sep string, tokens ...string) string
	// variadic functions

	// func prototype
	// func Double(n int) ( result int)

	fmt.Println(Sum(1, 2, 3, 1, 41, 23, 4124))
	fmt.Println(Concat(" \t", "askjd", "askdjashkdhsa", "1223"))

	var f func(n int) int

	f = Double
	g := Apply

	fmt.Println(f(6))
	fmt.Println(g(6, f))

	isMultipleOfX := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}

	var isMultipleOf3 = isMultipleOfX(3)
	var isMultipleOf9 = isMultipleOfX(9)

	fmt.Println(isMultipleOf3(6))
	fmt.Println(isMultipleOf9(89))
	fmt.Println(isMultipleOf3(91))

	//ranging over func
	func(yield func() bool)
	func(yield func(V) bool)
	func(yield func(K, V) bool)

}
