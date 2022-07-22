package solid

import "fmt"

// LSP: Liskov Substitutive Principle
// Doesn't really apply to Go, as there is no inheritance here
// The instructor will present an example that sports a variant of the LSP and is applicable to Go

// For this example, we create the Sized interface for geometrical shapes
type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

// And define our first shape as a rectangle
type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

// Everything seems to work for the rectangle, so let's create a Square
// Squares have width and height too, so we can say it has the same attributes as the Rectangle
type Square struct {
	Rectangle
}

// The difference is that a Square has equal width and height
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size

	return &sq
}

func (s *Square) GetWidth() int {
	return s.width
}

// And whenever we set the width, the height is set to the same value
func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) GetHeight() int {
	return s.height
}

// Or vice-versa
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// If this function is supposed to work with all shapes
func UseIt(sized Sized) {
	width := sized.GetWidth()

	// This statement here breaks the assumptions we make on how the square works
	// Because we're expecting to set only the height, not the width
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected an area of", expectedArea, "but got", actualArea)
}

// The behaviour of implementors of a particular type, like in this case, the Sized interface
// Should not break the core fundamental behaviors we rely on (like the square being a square)
// What we could do is perhaps create another type

type Square2 struct {
	side int
}

// There is no right way, however, to solve this in a general way
func (s *Square2) Rectangle() Rectangle {
	return Rectangle{
		width:  s.side,
		height: s.side,
	}
}

func Lsp() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)

	sq2 := Square2{5}
	sq2Rec := sq2.Rectangle()
	UseIt(&sq2Rec)
}
