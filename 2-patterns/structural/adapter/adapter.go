package adapter

import (
	"fmt"
	"strings"
)

// IRL Example:
// Electrical devices have different power (interface) requirements
//		Voltage (5V, 220V)
//		Socket/plug type (Europe, UK, USA)

// We cannot modify our gadgets to support every possible interface
//		Some support possible (e.g., 120/220V)

// Thus, we use a special device (an adapter) to give us the interface we require from the interface we have

// ADAPTER - A CONSTRUCT WHICH ADAPTS AN EXISTING INTERFACE X TO CONFORM TO THE REQUIRED INTERFACE Y

// Imagine a vectorial plane example
type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{
		[]Line{
			{0, 0, width, 0},
			{0, 0, 0, height},
			{width, 0, width, height},
			{0, height, width, height},
		},
	}
}

// Above, there's the interface I'm given

// Let's imagine we don't have any way of returning graphical elements

// Below, let's see the interface we have:

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}

	maxX += 1
	maxY += 1

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// Our problem is:
// the only way to create a rectangle is by creating a *VectorImage
// and the only way to draw something is by providing a RasterImage

// Solution:
// Build a new type that implements the desired interface
type vectorToRasterAdapter struct {
	points []Point
}

// Add a method that converts the provided type to the desired interface
func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("we have", len(a.points), "points")
}

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func (v vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func Adapter() {
	rc := NewRectangle(6, 4)
	// fmt.Println(DrawPoints(rc)) - can't print the rectangle
	adapter := VectorToRaster(rc)
	_ = VectorToRaster(rc) // unnecessary 2nd adapter call
	fmt.Println(DrawPoints(adapter))
}
