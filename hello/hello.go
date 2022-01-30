package main

import (
	"fmt"

	"moohoo/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("T0ddi")
	fmt.Println(message)
}
