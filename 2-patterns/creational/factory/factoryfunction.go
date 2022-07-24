package factory

import "fmt"

// Object creation logic may become too complicated
// Inner structs may have too many fields that all need to be initialized correctly
// Wholesale object creation (non-piecewise, unlike Builder) can be outsourced to
//		A separate function (Factory Function, a.k.a. Constructor)
//		A different, separate struct (Factory)

// FACTORY - A COMPONENT RESPONSIBLE SOLELY FOR THE WHOLESALE (NOT PIECEWISE) CREATION OF OBJECTS

// 1. The regular use of constructor goes like this... We create our struct
type Person struct {
	Name string
	Age  int

	// 3. EyeCount is an obvious attribute that is almost always 2
	// 4. We don't want it to receive other values than 2
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age, 2}
}

func FactoryFunction() {
	// 2. And then we construct it with with the curly braces {}
	p := Person{"John", 22, 2}

	// 5. Using the new function, we can create a Person with all 3 attribute, but we only really pass 2 to the function
	p2 := NewPerson("Jane", 24)

	fmt.Printf("p: %v\n", p)
	fmt.Printf("p2: %v\n", p2)
}
