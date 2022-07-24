package factory

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// Functional Approach
// How does this work?
// We have this function NewEmployeeFactory, that takes a position and an annualIncome
// Rather than returning an employee, it returns a an EmployeeFactoryFunction
// 	Everytime this EmployeeFactoryFunction receives a name, it will create a new Employee that
//		Has the same position and annualIncome as all other for that factory
func NewFunctionalEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// Structural Approach
// Works similarly to the Function Approach
// Is slightly more flexible, however, as you can keep attributes public if you need to tweak them later
// Once a Functional Factory is created, you can't tweak its values
// The issue with Structural Factories is that their methods need to be known by the API user
// It may be better, then, to have an Interface Factory/Contract, whose methods are all known
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func NewStructuralEmployeeFactory(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func FactoryGenerator() {
	developerFactory := NewFunctionalEmployeeFactory("Developer", 60000)
	managerFactory := NewFunctionalEmployeeFactory("Manager", 80000)

	adam := developerFactory("Adam")
	jane := managerFactory("Jane")

	fmt.Printf("developer: %v\n", adam)
	fmt.Printf("manager: %v\n", jane)

	bossFactory := NewStructuralEmployeeFactory("CEO", 100000)
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Amanda")
	fmt.Printf("boss: %v\n", boss)
}
