package main

import (
	"fmt"

	"example/greetings"

	"log"
)

func main() {

	names := []string{"Angel", "Samantha", "Dariana"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
