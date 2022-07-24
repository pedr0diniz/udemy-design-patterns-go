package factory

import "fmt"

type Human interface {
	SayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry, I'm too tired to talk")
}

// We can have our factory method return an interface rather than the concrete types
// This way, not only do we encapsulate the attributes or the inner concrete objects
// But we also allow the generate objects to have different behaviors
func NewHuman(name string, age int) Human {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func InterfaceFactory() {
	p := NewHuman("James", 34)
	p.SayHello()

	p2 := NewHuman("Eleanor", 101)
	p2.SayHello()
}
