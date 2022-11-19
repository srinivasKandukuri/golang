package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("breathe")
}

func (d dog) walk() {
	fmt.Println("can walk")
}

type pet1 struct {
	animal
	name string
}

func main() {
	p := dog{age: 20}

	p1 := pet1{name: "Jammy", animal: p}

	fmt.Println(p1.name)
	p1.breathe()

	p1.walk()

}
