package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(Name string) *person {
	p := person{name: Name}
	p.age = 13
	return &p
}
func main() {
	fmt.Println(person{"John", 42})
	fmt.Println(person{name: "Andy", age: 13})
	fmt.Println(person{age: 13})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Lucas"))

	s := person{"Cory", 13}
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 14
	fmt.Println(sp.age, s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		false,
	}
	fmt.Println(dog)
}
