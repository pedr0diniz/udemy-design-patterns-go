package main

import (
	"fmt"

	"github.com/pedr0diniz/2-patterns/creational/builder"
)

func main() {
	fmt.Println("Builder Pattern Basics:")
	builder.BuilderPattern()

	fmt.Println("\nBuilder Facets:")
	builder.BuilderFacets()

	fmt.Println("\nBuilder Parameters:")
	builder.BuilderParameter()

	fmt.Println("\nFunctional Builders:")
	builder.FunctionalBuilder()
}
