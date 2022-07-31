package main

import (
	"fmt"

	"github.com/pedr0diniz/2-patterns/structural/decorator"
)

func main() {
	fmt.Println("Multiple Aggregation:")
	decorator.MultipleAggregation()

	fmt.Println("\nDecorator:")
	decorator.Decorator()
}
