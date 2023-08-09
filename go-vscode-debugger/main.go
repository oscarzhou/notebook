package main

import (
	"fmt"
	"os"
)

func main() {
	foo := os.Getenv("foo")
	fmt.Printf("foo: %s\n", foo)

	message := "Greetings"
	target := "Earthlings"
	fmt.Printf("%s, %s!\n", message, target)

	// Change the value of the message variable
	message = "Salutations"
	fmt.Printf("%s, %s!\n", message, target)

	// Change the value of the target variable
	target = "Martians"
	fmt.Printf("%s, %s!\n", message, target)
}
