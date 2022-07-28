package main

import (
	"fmt"

	"github.com/pedr0diniz/2-patterns/structural/adapter"
)

func main() {
	fmt.Println("Adapter:")
	adapter.Adapter()

	fmt.Println("\nAdapter Caching:")
	adapter.AdapterCaching()
}
