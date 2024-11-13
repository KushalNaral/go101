package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, error := greetings.Hello("Padwan")

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	if error != nil {
		log.Fatal(error)
	}

	names := []string{"adam", "john", "mary"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	fmt.Println(messages)
}
