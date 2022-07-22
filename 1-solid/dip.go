package solid

import "fmt"

// DIP: Dependency Inversion Principle
// High Level Modules should not depend on Low Level Modules
// Both kinds of modules should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low-level module (Data storage, closer to hardware modules)
type Relationships struct {
	// What happens if the Relationships struct changes? It breaks our Research struct.
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := []*Person{}

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module (Business Logic)
type Research struct {
	// break DIP by depending on low-level module
	relationships Relationships
}

func (r *Research) Investigate() {
	// With this approach, we're directly accessing a field in the relationships
	// If the relationships struct changes, this breaks
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called", rel.to.name)
		}
	}
}

type DIPResearch struct {
	browser RelationshipBrowser
}

func (dr *DIPResearch) Investigate() {
	// With this approach, however, we're getting our results from a method in the lower-level module
	// We don't need to know the implementation details or access any property directly
	// All we need is to receive a []*Person fom the low-level module, no matter how
	// Also, the low-level module is now an abstraction (interface), rather than a concrete implementation
	for _, p := range dr.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func Dip() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{relationships}
	r.Investigate()

	dr := DIPResearch{&relationships}
	dr.Investigate()
}
