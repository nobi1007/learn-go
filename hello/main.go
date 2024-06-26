package main

import (
	"fmt"

	"github.com/nobi1007/learn-go/greetings"
)

func main() {
	msg := greetings.Hello("Shyam")
	fmt.Println(msg)
}
