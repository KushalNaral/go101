package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

const MaxRand = 15

var wg sync.WaitGroup

func StatRandomNumbers(n int) (int, int) {

	var a, b int
	for i := 0; i < n; i++ {
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++
		}
	}

	return a, b
}

func foo() {
	return
}

func foo2() string {
	return "bar"
}

func CompareLower4Bits(m, n uint32) (r bool) {
	r = m&0xF > n&0xF
	return true
}

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2

		time.Sleep(d)
	}

	wg.Done()
}

func Triple(n int) (r int) {
	defer func() {
		r += n
	}()

	return n + n
}

// func main() {
//
// 	fmt.Println("hello")
// 	defer func() {
// 		recover()
// 		fmt.Println("reciovered ehre")
// 	}()
//
// 	go func() {
// 		time.Sleep(time.Second)
// 		panic(123)
// 	}()
//
// 	for {
// 		time.Sleep(time.Second)
// 	}
//
// }
