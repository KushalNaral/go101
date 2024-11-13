package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {

	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func mainGenercs() {

	//simple composite literal
	// [3][4]int

	matrix := [3][4]int{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
	}
	fmt.Println(matrix)

	matrix[1][1] = 3

	fmt.Println(matrix)

	a := matrix[1]

	fmt.Println(a)

	//chan *[16]byte

	// c := make(chan *[16]byte)
	//
	// go func() {
	//
	// 	var dataA, dataB = new([16]byte), new([16]byte)
	// 	for {
	// 		_, err := rand.Read(dataA[:])
	// 		if err != nil {
	// 			close(c)
	// 		} else {
	// 			c <- dataA
	// 			dataA, dataB = dataB, dataA
	// 		}
	// 	}
	// }()
	//
	// for data := range c {
	// 	fmt.Println((*data)[:])
	// 	time.Sleep(time.Second / 2)
	// }

	addone := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }
	double := func(x int) int { return x * 2 }

	transforms := map[string][]func(int) int{
		"inc, inc, inc": {addone, addone, addone},
		"sqr, inc, dbl": {square, addone, double},
		"dbl, sqr, sqr": {double, square, square},
	}

	for _, n := range []int{2, 3, 5, 7} {
		fmt.Println(">>>", n)
		for name, transfers := range transforms {
			result := n
			for _, xfer := range transfers {
				result = xfer(result)
			}
			fmt.Printf(" %v: %v \n ", name, result)
		}
	}

	//Complex type
	// cT := []map[struct {
	// 	a int
	// 	b struct {
	// 		x string
	// 		y bool
	// 	}
	// }]interface {
	// 	Build([]byte, struct {
	// 		x string
	// 		y bool
	// 	}) error
	// 	Update(dt float64)
	// 	Destroy()
	// }

	type B = struct {
		x string
		y bool
	}

	type K = struct {
		a int
		b B
	}

	type E = interface {
		Build([]byte, B) error
		Update(dt float64)
		Destroy()
	}

	type T = []map[K]E

	var x T

	fmt.Printf("\n%V\n", x)

	aa, bb := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-aa, <-bb))
}
