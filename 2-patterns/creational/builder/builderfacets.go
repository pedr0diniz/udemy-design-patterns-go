package builder

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	CompanyName, Position string
	AnnualIncome          int
}

// We can effectively have one builder for all the parameters above, but
// It would probably make more sense for us to have an address builder and a job builder, wouldn't it?
type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	// PersonAddressBuilder points to PersonBuilder, which points to the Person object we're currently building.
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	// PersonJobBuilder points to PersonBuilder, which points to the Person object we're currently building.
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		&Person{},
	}
}

// PersonAddressBuilder has the same attributes a PersonBuilder has: a pointer to a person.
type PersonAddressBuilder struct {
	PersonBuilder
}

// That's why it accesses the person directly in its methods.
func (pab *PersonAddressBuilder) AtStreet(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) InCity(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostalCode(postalCode string) *PersonAddressBuilder {
	pab.person.Postcode = postalCode
	return pab
}

// PersonJobBuilder has the same attributes a PersonBuilder has: a pointer to a person.
type PersonJobBuilder struct {
	PersonBuilder
}

// That's why it accesses the person directly in its methods.
func (pjb *PersonJobBuilder) AtCompany(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func BuilderFacets() {
	pb := NewPersonBuilder()
	pb.
		Lives().AtStreet("123 London Road").InCity("London").WithPostalCode("SW12BC").
		Works().AtCompany("Fabrikam").AsA("Programmer").Earning(123000)
	person := pb.Build()

	fmt.Printf("person: %+v\n", person)

}
