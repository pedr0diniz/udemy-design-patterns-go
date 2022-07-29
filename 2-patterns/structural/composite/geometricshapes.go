package composite

import (
	"fmt"
	"strings"
)

// Objects use other objects' fields/methods through embedding
// Composition lets us make compound objects
//		E.g., a mathematical expression composed of simple expressions; or
//		A shape group made of several different shapes

// Composite design pattern is used to treat both single (scalar) and composite objects uniformly
// I.e., Foo and []Foo have common APIs

// COMPOSITE - A MECHANISM FOR TREATING INDIVIDUAL (SCALAR) OBJECTS AND COMPOSITIONS OF OBJECTS IN A UNIFORM MANNER.

// GraphicObject (with Name and Color) is our scalar object
// Whenever we add children, we apply the composite pattern
type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	// Recursiveness is key for dealing with possibly endless composite objects
	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{
		Name:     "Circle",
		Color:    color,
		Children: nil,
	}
}

func NewSquare(color string) *GraphicObject {
	return &GraphicObject{
		Name:     "Square",
		Color:    color,
		Children: nil,
	}
}

func GeometricShapes() {
	drawing := GraphicObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewSquare("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))
	drawing.Children = append(drawing.Children, group)

	fmt.Printf("drawing.String(): %v\n", drawing.String())
}
