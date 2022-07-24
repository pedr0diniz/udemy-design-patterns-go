package main

import (
	"fmt"

	"github.com/pedr0diniz/2-patterns/creational/factory"
)

func main() {
	fmt.Println("Factory Function:")
	factory.FactoryFunction()

	fmt.Println("\nInterface Factory Function:")
	factory.InterfaceFactory()

	fmt.Println("\nFactory Generators:")
	factory.FactoryGenerator()

	fmt.Println("\nPrototype Factory Function:")
	factory.PrototypeFactory()
}
