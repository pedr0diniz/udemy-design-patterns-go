package factory

import "fmt"

// Another thing we can have is a PrototypeFactory
// Prototypes are pre-built objects that don't necessarily contain all fields they need
type PrototypeEmployee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// In our case, we're creating prototype employees with pre-defined roles and salaries, but without a name
func NewPrototypeEmployee(role int) *PrototypeEmployee {
	switch role {
	case Developer:
		return &PrototypeEmployee{"", "Developer", 60000}
	case Manager:
		return &PrototypeEmployee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func PrototypeFactory() {
	// Our prototype factory method takes the role, returns our prototype employee
	m := NewPrototypeEmployee(Manager)

	// We update the prototype's name
	m.Name = "Pedro"

	// And it works just like any other factories
	fmt.Printf("m: %v\n", m)
}
