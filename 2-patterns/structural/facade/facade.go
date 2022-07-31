package facade

import "fmt"

// Balancing complexity and presentation/usability
// Imagine a home:
//		Many subsystems (electrical, sanitation)
//		Complex internal structure (e.g., floor layers)
//		End user is not exposed to internals
// Software is the same!
//		Many systems working to provide flexibility, but
//		API consumers want it to 'just work'

// FACADE - PROVIDES A SIMPLE, EASY TO UNDERSTAND USER INTERFACE OVER A LARGE AND SOPHISTICATED BODY OF CODE.

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width int, height int) *Buffer {
	return &Buffer{
		width:  width,
		height: height,
		buffer: make([]rune, width*height),
	}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

// Buffers can be hundreds of lines long
// We can create the Viewport in order to print in a better way
type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{
		buffer: buffer,
	}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// Facade struct
type Console struct {
	buffer    []*Buffer
	viewports []*Viewport
	offset    int
}

// Default scenario for a console with one buffer and one viewport
func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{
		buffer:    []*Buffer{b},
		viewports: []*Viewport{v},
		offset:    0,
	}
}

func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}

func Facade() {
	// Instead of working with low-level constructs, like buffers and viewports
	// We work directly with the console, abstracting the inner components logic
	c := NewConsole()

	// Does lots of things behind the scenes
	u := c.GetCharacterAt(1)
	fmt.Printf("u: %v\n", u)

	// This is the Facade: providing a simple API for something that's complicated
}
