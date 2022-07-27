package singleton

import "fmt"

// Singletons tend to break the dependency inversion principle

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
		// Depends directly on the singleton database, the concrete instance of the singleton database.
		// We should depend on abstractions rather than the concrete implementation.
	}
	return result
}

// Let's consider an unit testing scenario
func ProblemsWithSingleton() {
	cities := []string{"Seoul", "Mexico City"}

	// We're touching the "database" here. If we only want to test this function, why are we starting the whole system and its database connection?
	tp := GetTotalPopulation(cities)

	// If the population values from any of these cities change, this assertion breaks
	ok := tp == (17500000 + 17400000)
	fmt.Println("Populations of Seoul and Mexico City equal 34,900,000?", ok)
}
