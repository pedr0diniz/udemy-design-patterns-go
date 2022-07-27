package main

import (
	"fmt"

	"github.com/pedr0diniz/2-patterns/creational/prototype"
)

func main() {
	fmt.Println("Deep Copying:")
	prototype.DeepCopying()

	fmt.Println("\nCopy Method:")
	prototype.CopyMethod()

	fmt.Println("\nCopy Through Serialization:")
	prototype.CopyThroughSerialization()

	fmt.Println("\nPrototype Factory:")
	prototype.PrototypeFactory()
}
