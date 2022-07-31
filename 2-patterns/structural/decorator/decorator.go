package decorator

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c Circle) Render() string {
	return fmt.Sprintf("Circle of radius %f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s Square) Render() string {
	return fmt.Sprintf("Square with side %f", s.Side)
}

// What if we want to color these shapes now?
// We could add a "Color" field to the existing shapes, but this breaks the OCP
// Let's aggregate these types, then!

type ColoredShape struct {
	Shape Shape
	Color string
}

func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s has the color %s", c.Shape.Render(), c.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s has %f%% transparency", t.Shape.Render(), t.Transparency*100.0)
}

func Decorator() {
	circle := Circle{2}
	fmt.Printf("circle.Render(): %v\n", circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Printf("redCircle.Render(): %v\n", redCircle.Render())

	// One thing to remember is that, in this case, we cannot Resize the redCircle, because its inner type isn't a Circle
	// Its inner type is a Shape
	// And not all shapes have the Resize() method

	redHalfTransparentCircle := TransparentShape{&redCircle, 0.5}
	fmt.Printf("redHalfTransparentCircle.Render(): %v\n", redHalfTransparentCircle.Render())

	// The transparentShape shows that decorators can be chained
}
