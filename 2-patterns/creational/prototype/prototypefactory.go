package prototype

import (
	"encoding/json"
	"fmt"
)

type PFAddress struct {
	Suite               int
	StreetAddress, City string
}

type PFEmployee struct {
	Name   string
	Office PFAddress
}

func (p *PFEmployee) deepCopyThroughSerialization() *PFEmployee {
	// When we marshal/encode/serialize our objects, it becomes a slice of bytes.
	e, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding bytes to PersonWithFriends!")
	}

	// We can create a new object and unmarshal/decode/deserialize the slice of encoded bytes into the new object.
	result := PFEmployee{}
	if err = json.Unmarshal(e, &result); err != nil {
		fmt.Println("Error decoding bytes to PersonWithFriends!")
	}

	// The new object, then, is returned with the same values of the original one, but a new memory address.
	return &result
}

var mainOfficePrototypeEmployee = PFEmployee{
	Name: "",
	Office: PFAddress{
		Suite:         0,
		StreetAddress: "123 East Dr",
		City:          "London",
	},
}

var auxiliaryOfficePrototypeEmployee = PFEmployee{
	Name: "",
	Office: PFAddress{
		Suite:         0,
		StreetAddress: "666 West Dr",
		City:          "London",
	},
}

// "Manual" Prototype Approach, taking an employee, name and suite number
func newPFEmployee(proto *PFEmployee, name string, suite int) *PFEmployee {
	result := proto.deepCopyThroughSerialization()
	result.Name = name
	result.Office.Suite = suite

	return result
}

// Prototype Factory that already pre-defines the Main Office Address
func NewMainOfficePFEmployee(name string, suite int) *PFEmployee {
	return newPFEmployee(&mainOfficePrototypeEmployee, name, suite)
}

// Prototype Factory that already pre-defines the Auxiliary Office Address
func NewAuxiliaryOfficePFEmployee(name string, suite int) *PFEmployee {
	return newPFEmployee(&auxiliaryOfficePrototypeEmployee, name, suite)
}

func PrototypeFactory() {
	// While it may seem harder to setup, it also allows for much easier object creation
	john := NewMainOfficePFEmployee("John", 100)
	jane := NewAuxiliaryOfficePFEmployee("Jane", 200)

	fmt.Println("john:", john, john.Office)
	fmt.Println("jane:", jane, jane.Office)
}
