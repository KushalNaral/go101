package main

import "fmt"

type _string struct {
	elements *byte
	len      int
}

func mainStrings() {

	const World = "world"
	var hello = "hi"

	var hw = hello + " " + World

	hello += World

	fmt.Println(hw)
	fmt.Println(len(hello))

	fmt.Println(hw[:1])

	var str = "world"
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	key := []byte{'k', 'e', 'y'}

	m := map[string]string{}

	m[string(key)] = "value"

	fmt.Println(m[string(key)])
}
