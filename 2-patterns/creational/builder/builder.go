package builder

import (
	"fmt"
	"strings"
)

// Simple objects can be built using a constructor
// Large objects require a lot of cerimony to create
// Having a constructor resolving 10 parameters isn't optional
// The Builder method allows piece-by-piece construction
// It does so by providing an API for step-by-step construction

// BUILDER - WHEN PIECEWISE OBJECT CONSTRUCTION IS COMPLICATED, PROVIDE AN API FOR DOING IT SUCCINTLY

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

// Fluent method calls allow you to chain calls rather by returning the receiver
func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)

	return b
}

func BuilderPattern() {
	// A built-in example is Go is the Strings builder:

	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")

	fmt.Printf("Printing from Go's string builder - sb.String(): %v\n", sb.String())

	words := []string{"hello", "world"}

	// If we want to re-use a builder, we must reset it before we start building something new
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")

	fmt.Printf("Printing from Go's string builder - sb.String(): %v\n", sb.String())

	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Printf("Printing from our HTML builder - b.String(): \n%v\n", b.String())

	bf := NewHtmlBuilder("ul")
	bf.AddChildFluent("li", "hello").AddChildFluent("li", "world")

	fmt.Printf("Printing from our HTML builder fluent calls - bf.String(): \n%v\n", bf.String())
}
