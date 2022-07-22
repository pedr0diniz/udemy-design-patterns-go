package solid

import "fmt"

// OCP: Open-Closed Principle
// Open for Extension, Closed for Modification

// We'll also learn an Enterprise Pattern called "Specification" here.

type Product struct {
	name  string
	color Color
	size  Size
}

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Filter struct{}

// Adding one method here is ok.
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := []*Product{}

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Adding another one means we come back to this file and modify what has already been tested and working
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := []*Product{}

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

// And yet another function addition just keeps making this more complex
// Wouldn't it be nice if we could add more filters without having to modify the place that has the current ones?
// That's what the OCP is about
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := []*Product{}

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Let's give the Specification Pattern a try
// First, we create the Specification interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

// Then, for each of our previous filters, we create a Specification struct
type ColorSpecification struct {
	color Color
}

// And each struct gets its own IsSatisfied method, checking if its attribute equals the product's attribute
// Just like the previous filters worked
func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

// This specification gives us flexibility to merge whichever two specifications
type AndSpecification struct {
	first, second Specification
}

// To satisfy the composite specification, we just check if both first and second specifications are satisfied
func (as AndSpecification) IsSatisfied(p *Product) bool {
	return as.first.IsSatisfied(p) && as.second.IsSatisfied(p)
}

// And now, we create our "better filter". This will ensure our system's flexibility
type BetterFilter struct{}

func (bf *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := []*Product{}
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func Ocp() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	// Old filters (creating a new function for every new filter)
	fmt.Printf("Green products (old): \n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// New filters
	fmt.Printf("\nGreen products (new): \n")

	// Instead of filtering directly by color, we now create a specification for that color.
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}

	// And where we used to pass the color as parameter, we now pass the color specification.
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// Well... ok, it's basically the same thing. How is using specifications better?
	// Because we can write new specifications without touching old ones
	// And the BetterFilter.Filter() method will work for any Specification that implements the interface
	// Let's check how we could filter by size:
	fmt.Printf("\nLarge products (new): \n")
	largeSpec := SizeSpecification{large}
	for _, v := range bf.Filter(products, largeSpec) {
		fmt.Printf(" - %s is large\n", v.name)
	}

	// Lastly, let's try the AndSpecification in order to see it replace our old "ColorAndSize" filter.
	largeGreenSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("\nLarge green products (new): \n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}

// The idea here is to not keep jumping to the same file and keep modifying it over and over again
