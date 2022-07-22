package main

import (
	"fmt"

	solid "github.com/pedr0diniz/1-solid"
)

func main() {
	fmt.Println("Single Responsibility Principle:")
	solid.Srp()

	fmt.Println("\nOpen-Closed Principle:")
	solid.Ocp()

	fmt.Println("\nLiskov Substitution Principle:")
	solid.Lsp()

	fmt.Println("\nInterface Segregation Principle:")
	solid.Isp()

	fmt.Println("\nDependency Inversion Principle:")
	solid.Dip()
}
