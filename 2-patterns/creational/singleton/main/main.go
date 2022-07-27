package main

import (
	"fmt"

	"github.com/pedr0diniz/2-patterns/creational/singleton"
)

func main() {
	fmt.Println("Singleton:")
	singleton.Singleton()

	fmt.Println("\nProblems with Singleton:")
	singleton.ProblemsWithSingleton()

	fmt.Println("\nSingleton and Dependency Inversion:")
	singleton.SingletonAndDependencyInversion()
}
