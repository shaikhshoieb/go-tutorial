package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"shoieb", 24})
	fmt.Println(person{name: "shoieb 1", age: 25})
	fmt.Println(person{name: "shoieb 2"})
	fmt.Println(&person{name: "shoieb 3", age: 27})
	fmt.Println(newPerson("shoieb 4"))

	sh := person{"shoieb 5", 28}
	fmt.Println(sh.name)

	sp := &sh
	fmt.Println(sp.name)

	sp.age = 30

	fmt.Println(sp.age)
	fmt.Println(sh.age)
	sh.age = 31
	fmt.Println(sp.age)
	fmt.Println(sh.age)

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "dog",
		isGood: true,
	}
	fmt.Println(dog.name)
	fmt.Println(dog.isGood)
	fmt.Println(dog)
}
