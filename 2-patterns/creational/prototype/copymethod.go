package prototype

import "fmt"

// This is a DeepCopy() method.
// It takes a pointer to a struct and returns a new memory address with the copied data
func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type PersonWithFriends struct {
	Person
	Friends []string
}

// By making our Person struct more complex, we now call the Address DeepCopy() method inside our PersonWithFriends DeepCopy()
func (p *PersonWithFriends) DeepCopy() *PersonWithFriends {

	// q is our new PersonWithFriends
	// It receives a pointer to the function caller
	q := *p

	// With a pointer to the function caller, we can deep copy its attributes
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends) // This copies slice values to a new slice

	// Since we're working with pointers and addresses, we return the address to the new PersonWithFriends
	return &q
}

func CopyMethod() {
	john := PersonWithFriends{Person{"John", &Address{"123 London Rd", "London", "UK"}}, []string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
