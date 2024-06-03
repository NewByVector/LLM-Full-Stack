package main

type Person struct {
	Name string
	Age  int
}

// 自动引用与解引用
func (p *Person) changeName(name string) {
	p.Name = name
}

// slice、map 是引用传递
func changeSlice(s []string) {
	s[0] = "x"
}

func main() {
	p := Person{Name: "John", Age: 30}
	p.changeName("Doe")
	println(p.Name)

	s := []string{"a", "b", "c"}
	changeSlice(s)
	println(s[0])
}
