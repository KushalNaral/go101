package main

import (
	"fmt"
	"reflect"
	"sync"
)

type Locakable[T any] struct {
	mu   sync.Mutex
	data T
}

type X[T any, Y string] struct {
	Data T
	string
}

func (l *Locakable[T]) Do(f func(*T)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	f(&l.data)
}

func NoDiff[V comparable](vs ...V) bool {

	if len(vs) == 0 {
		return true
	}

	v := vs[0]

	for _, x := range vs[1:] {
		if v != x {
			return false
		}
	}

	return true
}

type L interface {
	Run() error
	Stop()
}

type M interface {
	L
	Step() error
}

type N interface {
	M
	interface{ Resume() }
	~map[int]bool
	~[]byte | string
}

type O interface {
	Pause()
	N
	string
	int64 | ~chan int | any
}

type Bytes []byte
type Letters Bytes
type Blank struct{}
type MyString string

func (Bytes) M() {
}

func (Blank) M() {
}
func (MyString) M() {
}

type P interface{ []byte }

type Q interface{ ~[]byte }

type R interface{ []int32 | MyString }

type S interface {
	R
	M()
}

type T interface {
	~[]byte | ~string
}

type U interface {
	M()
}

type V interface {
	[]byte
	any
}

type W interface {
	T
	U
}

type C interface {
	comparable
	[]byte | func() | map[int]bool | string | [2]any
}

type A [2]any

func (a *A) M() {
}

type B struct {
	A
	x any
}

type CA interface {
	comparable
	M()
}

type Data[A int64 | int32, B byte | bool, C comparable] struct {
	a A
	b B
	c C
}

var x = Data[int64, byte, [8]byte]{1 << 62, 255, [8]byte{}}

type Y = Data[int32, bool, string]

type Z = Data[int64, uint8, [8]uint8]

type WZ Data[int64, byte, [8]byte]

type Ordered interface {
	~int | ~uint | ~int8 |
		~uint8 | ~int16 |
		~uint16 | ~int32 |
		~uint32 | ~int64 |
		~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Max[S ~[]E, E Ordered](vs S) E {

	if len(vs) == 0 {
		panic("no elements")
	}

	var r = vs[0]

	for i := range vs[1:] {
		if vs[i] > r {
			r = vs[i]
		}
	}

	return r
}

func Check[T ~string | ~uint32 | any](a T, b T, c T) T {
	println("a print")
	println(a)
	return a
}

//type Age int

var ages = []Age{1, 2, 3, 4, 51, 12, 41, 3}

var langs = []string{"Asda", "asdasd", "asdasdasdas", "Go", "1123123", "&asodjasdl", "", "  "}

type (
	Name = string
	AgeX = int
)

type table = map[string]int
type Table = map[Name]AgeX

// type Book struct {
// 	author string
// 	title  string
// 	pages  int
// }

func mainT() {

	println(reflect.TypeOf(x) == reflect.TypeOf(Z{}))
	println(reflect.TypeOf(x) == reflect.TypeOf(Y{}))

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", Z{})
	fmt.Printf("%T\n", Y{})

	var maxAge = Max[[]Age]
	println(maxAge(ages))

	var maxStr = Max[[]string]
	println(maxStr(langs))

	var check = Check[string]
	check("a", "b", "c")

	var check32 = Check[uint32]
	check32(1, 2, 3)

	checkAny := Check[bool]
	checkAny(true, false, false)

	// *int
	// **int

	type Ptr *int
	type PP *Ptr

}
