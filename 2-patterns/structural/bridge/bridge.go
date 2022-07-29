package bridge

import "fmt"

// Prevents a 'Cartesian Product' complexity explosion
// Example:
//		Common type ThreadScheduler
//		Can be preemptive
//		Can run on Windows or Unix
//		End up with a 2x2 scenario: WindowsPTS, UnixPTS, WindowsCTS, UnixCTS

// Bridge pattern avoids the entity "explosion" (multiplication of possibilities)

// Before:
// ThreadScheduler
//		PreemptiveThreadScheduler
//			WindowsPTS
//			UnixPTS
//		CooperativeThreadScheduler
//			WindowsPTS
//			UnixPTS

// After:
// PreemptiveThreadShceduler
//		ThreadScheduler
//			PlatformScheduler
// CooperativeThreadScheduler
// 		ThreadScheduler
//			PlatformScheduler
//
// Where a PlatformScheduler can be:
//		UnixScheduler
//		WindowsScheduler
//
// This way, we create less individual components

// BRIDGE - A MECHANISM THAT DECOUPLES AN INTERFACE (HIERARCHY) FROM AN IMPLEMENTATION (HIERARCHY).

// Circle, square shapes
// Raster, vector rendering

// RasterCicle, VectorCicle, RasterSquare, VectorSquare, ...

// By having the renderer interface in our shapes, we can define their behaviors based on which renderer is passed during the object construction
type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(side float32)
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (v *VectorRenderer) RenderSquare(side float32) {
	fmt.Println("Drawing a square of side", side)
}

type RasterRenderer struct{}

func (r RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

func (v *RasterRenderer) RenderSquare(side float32) {
	fmt.Println("Drawing pixels for a square of side", side)
}

type Square struct {
	renderer Renderer
	side     float32
}

func NewSquare(renderer Renderer, side float32) *Square {
	return &Square{
		renderer: renderer,
		side:     side,
	}
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.side)
}

func (s *Square) Resize(factor float32) {
	s.side *= factor
}

type Circle struct {
	renderer Renderer
	radius   float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{
		renderer: renderer,
		radius:   radius,
	}
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

// If we wanted another shape, we wouldn't need to create two new objects, VectorShape and RasterShape
// We would only need to create the one new shape and make it receive a renderer, as done by the Circle and the Square

func Bridge() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	fmt.Println("Raster Circle")
	circle1 := NewCircle(&raster, 5)
	circle1.Draw()
	circle1.Resize(2)
	circle1.Draw()

	fmt.Println("\nRaster Square")
	square1 := NewSquare(&raster, 5)
	square1.Draw()
	square1.Resize(2)
	square1.Draw()

	fmt.Println("\nVector Circle")
	circle2 := NewCircle(&vector, 5)
	circle2.Draw()
	circle2.Resize(2)
	circle2.Draw()

	fmt.Println("\nVector Square")
	square2 := NewSquare(&vector, 5)
	square2.Draw()
	square2.Resize(2)
	square2.Draw()
}
