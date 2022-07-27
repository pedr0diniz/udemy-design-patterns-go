package singleton

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// For some components it only makes sense to have one in the system, such as
//		Database repository
//		Object factory

// For these objects, the construction call is either expensive or is unnecessary to do more than once
//		We only do it once
// 		We give everyone the same instance

// We want to prevent anyone to create additional copies
// Need to take care of lazy instantiation

// SINGLETON - A COMPONENT THAT IS INSTANTIATED ONLY ONCE

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// Imported from the instructor's materials
func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

// Go's standard package type for defining singletons, ensuring something is only called once
var once sync.Once

// We want thread safety here
// We don't want two threads initializing this object at the same time
// To ensure thread safety, we have two options:
//		1. sync.Once init()
//		2. Laziness (you only construct this database whenever someone asks for it)

var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	// Runs its parameters function only once
	once.Do(func() {
		caps, e := readData("2-patterns/creational/singleton/capitals.txt")
		db := singletonDatabase{caps}
		if e == nil {
			db.capitals = caps
		}
		instance = &db
	})

	// Always returns the same instance
	return instance
}

func Singleton() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Printf("Population of Seoul: %v\n", pop)
}
