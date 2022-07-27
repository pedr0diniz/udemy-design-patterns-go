package singleton

import "fmt"

type Database interface {
	GetPopulation(name string) int
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func SingletonAndDependencyInversion() {
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println("Populations of Alpha and Gamma equal 4?", tp == 4)
}

// The Singleton itself isn't the problem
// The problem is depending directly on the Singleton
// Applying the Dependency Inversion solves this problem
