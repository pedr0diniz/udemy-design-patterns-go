package solid

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// SRP: Single Responsibility Principle
// This principle carries the idea that classes/structs should only have one responsibility

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)

	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

// we need to apply the SEPARATION OF CONCERNS
// if we don't apply the separation of concerns, we create god objects

// the two functions below break the separation of concerns.
// it is the Journal's concern to deal with its entries
// not to deal with persistence.
func (j *Journal) Save(filename string) {
	// ...
}

func (j *Journal) SaveToFile(filename string) {
	// ...
}

// it is better to separate the persistence to another struct instead
// with a separated struct, we can create methods that make sense
// for its responsibilities
type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func Srp() {
	// Instantiating a journal and handling its entries
	j := Journal{}
	j.AddEntry("I'm pissed today")
	j.AddEntry("My belly hurts")
	fmt.Println(j.String())

	// Handling persistence
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
