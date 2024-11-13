package main

import (
	"bytes"
	"fmt"
	"reflect"
)

const Size = 32

type Person struct {
	Name string
	Age  int
}

type language struct {
	name string
	year int
}

type z []int
type w [2]int
type y *w

func mainContainer() {

	var ll []int
	var ass = make([]int, 0)

	var x0 = (*[0]int)(ll)
	var y0 = (*[0]int)(ass)

	_, _ = x0, y0

	s1 := []int{1, 2, 3}
	s2 := s1

	a1 := [3]int{}
	a2 := a1

	a1[0] = 20

	s1[0] = 10
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(a1)
	fmt.Println(a2)

	// //array
	// [N]T
	// //slice
	// []T
	// //map
	// map[K]T

	// 	T = container type, only these can be stored in the container
	// N = non negatice int constant basically length
	// K = comparable type type of key of a map type

	s5 := [5]string{}

	m1 := map[string]int{}
	mbool := map[uint32]bool{}
	mint := map[int16][6]string{}
	mstr := map[struct{ x int }]*int8{}

	sArr := [500]Person{}

	sArr[0] = Person{
		Name: "asdsd",
		Age:  123,
	}

	mcompa := map[string]int{"C": 123, "p": 12}
	masso := map[int][][]int{
		1: {
			{1, 2, 3}, {5, 6, 7},
		},
	}

	nindex := [...]int{1, 2, 3}
	ncopy := nindex
	nindex[0] = 123123

	var heads = []*[4]byte{
		{'P', 'N', 'G', ' '},
	}

	var hD = []*[4]byte{}

	var lA = [...]language{
		language{
			"C", 19231,
		}, language{
			"Pausid", 12312,
		},
		language{
			"G", 1190,
		},
	}

	var l = []language{
		language{
			"C", 19231,
		}, language{
			"Pausid", 12312,
		},
		language{
			"G", 1190,
		},
	}

	type LangCategory struct {
		dynamic bool
		string  bool
	}

	var lC = map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"Py": 123,
			"C":  1001,
		},
		LangCategory{true, false}: map[string]int{
			"asjdkj": 123,
		},
		LangCategory{false, true}: map[string]int{
			"false": 123,
		},
		LangCategory{false, false}: map[string]int{
			"true": 123,
		},
	}

	lct := LangCategory{true, false}
	lcf := LangCategory{true, false}

	fmt.Println(s5)
	fmt.Println(m1)
	fmt.Println(mbool)
	fmt.Println(mint)
	fmt.Println(mstr)
	fmt.Println(sArr)
	fmt.Println(mcompa)
	fmt.Println(masso)
	fmt.Println(nindex)
	fmt.Println(ncopy)
	fmt.Println(heads)
	fmt.Println(hD)
	fmt.Println(reflect.TypeOf(l))
	fmt.Println(reflect.TypeOf(lA))
	fmt.Println(lC)
	fmt.Println(lct)
	fmt.Println(lcf)
	fmt.Println(lC[lct]["asjdkj"])
	fmt.Println(cap(sArr))

	m0 := map[int]int{1: 0, 2: 3, 3: 2, 4: 5, 5: 5}
	ma := m0

	ma[0] = 2
	fmt.Println(m0, ma)

	ms := map[int]int{1: 1, 2: 2, 3: 3}
	fmt.Println(ms)

	delete(ms, 1)
	fmt.Println(lC)
	delete(lC, lct)
	fmt.Println(lC)

	as := []int{1, 2, 3}
	as1 := append(as, 4, 5, 6, 7)

	fmt.Println(as)
	fmt.Println(as1)

	as[0] = 21

	fmt.Println(as)
	fmt.Println(as1)

	fmt.Println(cap(as))
	fmt.Println(cap(as1))

	fmt.Println(len(as))
	fmt.Println(len(as1))

	s21 := append(as1, []int{21, 21}...)
	fmt.Println((s21))
	fmt.Println(len(s21))
	fmt.Println(cap(s21))

	//make(M, n) // M -> map type, n -> entries count

	m := make(map[string]int)
	fmt.Println(m)

	mp := *new(map[string]int)
	fmt.Println(mp)

	a := [...]int{0, 1, 2, 6}
	// deriving slices : baseContainer[low : high : max]
	fmt.Println(a)
	s0 := a[:]
	fmt.Println(s0)

	a[0] = 12313

	fmt.Println(a)
	fmt.Println(s0)

	type Ta []int
	type Tb []int

	dest := Ta{1, 2, 3, 5}
	src := Tb{5, 6, 7, 8, 9}

	nn := copy(dest, src)
	fmt.Println(nn, dest)

	nn = copy(dest[1:], dest)
	fmt.Println(nn, dest)

	//iteration for key, elem = range container {}

	type Person struct {
		name string
		age  int
	}

	persons :=
		[2]Person{{"Alice", 12}, {"asdljs", 123}}

	for _, person := range persons {
		fmt.Println("name is ", person.name, "age is : ", person.age)
	}

	mdyna := map[int]struct{ dynamic, strong bool }{
		0: {true, false},
		1: {true, true},
		2: {false, true},
		3: {false, false},
	}

	fmt.Println(mdyna)

	for _, v := range mdyna {
		v.dynamic, v.strong = true, true
	}

}

type Buffer struct {
	start, end int
	data       [1024]byte
}

func fa(buffers []Buffer) int {
	numUnreads := 0

	for _, buf := range buffers {
		numUnreads += buf.end - buf.start
	}

	return numUnreads
}

func fb(buffers []Buffer) int {
	numUnreads := 0

	for i := range buffers {
		numUnreads += buffers[i].end - buffers[i].start
	}

	return numUnreads
}
