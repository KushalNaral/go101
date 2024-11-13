package main

import (
	"fmt"
)

type Ptr *int
type PP *Ptr

//pointer value stores the address of another value we can say the pointer value directyle references the other value

func double(x *int) {
	*x += *x
}

func newInt() *int {
	a := 3
	return &a
}

func mainPoint() {

	p0 := new(int)
	fmt.Println(p0)
	fmt.Println(*p0)

	x := *p0

	p1, p2 := &x, &x
	fmt.Println(p1 == p2)
	fmt.Println(p0 == p1)
	fmt.Println(p0 == p2)

	p3 := &(*p0)
	fmt.Println(p3)

	*p0, *p1 = 123, 12331
	fmt.Println(*p2, x, *p3)

	var a = 3
	double(&a)
	fmt.Println(a)

	var p = &a
	double(p)
	fmt.Println(*p)

	var y = newInt()
	var z = newInt()

	fmt.Println(y)
	fmt.Println(z)

	*z = 15

	fmt.Println(*y)
	fmt.Println(*z)

	ab := int64(5)
	pq := &ab

	*pq++

	fmt.Println(*pq, ab)

}
