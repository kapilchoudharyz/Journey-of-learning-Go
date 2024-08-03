package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("kapil")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
	message, error = greetings.Bye("")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}
