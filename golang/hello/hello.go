package main

import (
	"example/greetings"
	"fmt"
	"log"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"))

	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message, err := greetings.Hello("vector")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"vector", "alice", "bob"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
