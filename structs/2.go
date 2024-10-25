package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func initialize_person(name string, age int) *Person {
	person := Person{
		Name: name,
		Age:  age,
	}
	return &person
}
func main() {
	p := initialize_person("Ankush Banerjee", 22)
	fmt.Println("age:", p.Name)
	fmt.Println("Name", p.Age)

	p.Age = 31
	fmt.Println("Updated Age:", p)

}
