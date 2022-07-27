package prototype

import "fmt"

// Complicated objects aren't designed from scratch
//		They reiterate existing objects/designs
// An existing (partially or fully constructed) design is a Prototype
// We make a copy of the prototype and customize it
//		This requires "deep copy" support
//		Be careful with pointers when copying
// Makes cloning convenient (e.g., via a Factory)

// PROTOTYPE - A PARTIALLY OR FULLY INITIALIZED OBJECT THAT YOU COPY (CLONE) AND MAKE USE OF

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func DeepCopying() {
	john := Person{"John", &Address{"123 London Rd", "London", "UK"}}

	// Copies all pointers. Jane's address shares the same memory address as John's
	jane := john

	jane.Name = "Jane"

	// This, is a deep copy. It creates another object, but with pre-existing values.
	jane.Address = &Address{
		StreetAddress: john.Address.StreetAddress,
		City:          john.Address.City,
		Country:       john.Address.Country,
	}

	// If we didn't make the deep copy and left Jane's address as it was,
	// Changing it would change John's address as well.
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println("john:", john, john.Address)
	fmt.Println("jane:", jane, jane.Address)
}
