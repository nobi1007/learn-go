package main

import (
	"fmt"

	"github.com/nobi1007/learn-go/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("S")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
