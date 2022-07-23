package builder

import "fmt"

type PersonF struct {
	name, position string
}

type personMod func(*PersonF)

type PersonFBuilder struct {
	actions []personMod
}

func (b *PersonFBuilder) Called(name string) *PersonFBuilder {
	b.actions = append(b.actions, func(p *PersonF) {
		p.name = name
	})
	return b
}

func (b *PersonFBuilder) WorksAsA(position string) *PersonFBuilder {
	b.actions = append(b.actions, func(p *PersonF) {
		p.position = position
	})
	return b
}

func (b *PersonFBuilder) Build() *PersonF {
	p := PersonF{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// The functional builder makes our building process more easily extensible and also delayed
// Building steps aren't run on the go, but are stacked to be run all at once when the .Build() method is called
func FunctionalBuilder() {
	pfb := PersonFBuilder{}
	person := pfb.Called("Dmitri").WorksAsA("Developer").Build()
	fmt.Printf("person: %+v\n", person)
}
