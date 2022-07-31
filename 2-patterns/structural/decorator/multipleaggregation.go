package decorator

import "fmt"

// Want to augment an object with additional functionality
// Do not want to rewrite or alter existing code (OCP)
// Want to keep new functionality separate (SRP)
// We need to be able to interact with existing structures
// Solution: embed the decorated object and provide additional functionality

// DECORATOR - FACILITATES THE ADDITION OF BEHAVIORS TO INDIVIDUAL OBJECTS THROUGH EMBEDDING.

// In our example, let's imagine that we have a dragon that's both a Bird and a Lizard at the same time

// Getters and setters go against idiomatic Go, but there's no way around them in this case.
type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int       { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int       { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}

// Here, Dragon is a decorator
// It extends the behaviors of its inners types
// And kinda works as a proxy to them
type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int { return d.bird.age }
func (d *Dragon) SetAge(age int) {
	d.bird.age = age
	d.lizard.age = age
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{
		bird:   Bird{},
		lizard: Lizard{},
	}
}

func MultipleAggregation() {
	// Decorator approach:
	d := NewDragon()
	d.SetAge(11)
	d.Fly()
	d.Crawl()
}
