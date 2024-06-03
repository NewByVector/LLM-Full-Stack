package main

import "fmt"

func main() {
	m := map[string]string{"name": "John", "age": "30"}

	m["favorite"] = "golang"
	m["age"] = "32"
	delete(m, "name")

	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
}
