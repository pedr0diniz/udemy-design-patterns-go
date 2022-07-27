package prototype

import (
	"encoding/json"
	"fmt"
)

func (p *PersonWithFriends) deepCopyThroughSerialization() *PersonWithFriends {
	// When we marshal/encode/serialize our objects, it becomes a slice of bytes.
	e, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding bytes to PersonWithFriends!")
	}

	// We can create a new object and unmarshal/decode/deserialize the slice of encoded bytes into the new object.
	result := PersonWithFriends{}
	if err = json.Unmarshal(e, &result); err != nil {
		fmt.Println("Error decoding bytes to PersonWithFriends!")
	}

	// The new object, then, is returned with the same values of the original one, but a new memory address.
	return &result
}

func CopyThroughSerialization() {
	john := PersonWithFriends{Person{"John", &Address{"123 London Rd", "London", "UK"}}, []string{"Matt", "Chris"}}

	// Copies all pointers. Jane's address shares the same memory address as John's before the deep copy.
	jane := john.deepCopyThroughSerialization()

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
