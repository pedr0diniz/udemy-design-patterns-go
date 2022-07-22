package solid

import "fmt"

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// Does it all
type MultiFunctionPrinter struct {
}

func (mfp MultiFunctionPrinter) Print(d Document) {
	fmt.Println("I'm a MultiFunctionPrinter printing a document!")
}

func (mfp MultiFunctionPrinter) Fax(d Document) {
	fmt.Println("I'm a MultiFunctionPrinter faxing a document!")
}

func (mfp MultiFunctionPrinter) Scan(d Document) {
	fmt.Println("I'm a MultiFunctionPrinter scanning a document!")
}

// Cannot scan or fax
type OldFashionedPrinter struct {
}

func (ofp OldFashionedPrinter) Print(d Document) {
	fmt.Println("I'm an OldFashionedPrinter printing a document!")
}

// Deprecated
func (ofp OldFashionedPrinter) Fax(d Document) {
	fmt.Println("Error! I'm an OldFashionedPrinter, I can't fax!")
}

// Deprecated
func (ofp OldFashionedPrinter) Scan(d Document) {
	fmt.Println("Error! I'm an OldFashionedPrinter, I can't scan!")
}

// If we keep just one interface with all methods, we force the OldFashionedPrinter to implement features it doesn't have
// Therefore, it would be better to split the functionalities in different interfaces

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxxer interface {
	Fax(d Document)
}

// Not forced to implement anything it doesn't do
type RegularPrinter struct{}

func (rp RegularPrinter) Print(d Document) {
	fmt.Println("I'm a printer printing a document!")
}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	fmt.Println("I'm a Photocopier printing a document!")
}

func (p Photocopier) Scan(d Document) {
	fmt.Println("I'm a Photocopier scanning a document!")
}

// go allows you to build composite interfaces
// This can be an elegant solution
type MultiFunctionalDevice interface {
	Printer
	Scanner
}

// Another thing that you can do is implement a Decorator pattern
type MultiFunctionalMachine struct {
	printer Printer
	scanner Scanner
}

// And have interfaces as components to make use of polymorphism
func (mfm MultiFunctionalMachine) Print(d Document) {
	mfm.printer.Print(d)
}

func (mfm MultiFunctionalMachine) Scan(d Document) {
	mfm.scanner.Scan(d)
}

func Isp() {
	doc := Document{}

	mfp := MultiFunctionPrinter{}
	fmt.Println("What can the multifunction printer do?")
	mfp.Print(doc)
	mfp.Fax(doc)
	mfp.Scan(doc)

	ofp := OldFashionedPrinter{}
	fmt.Println("\nWhat can the old-fashioned printer do?")
	ofp.Print(doc)
	ofp.Fax(doc)
	ofp.Scan(doc)

	rp := RegularPrinter{}
	fmt.Println("\nWhat can the regular printer do?")
	rp.Print(doc)

	pc := Photocopier{}
	fmt.Println("\nWhat can the photocopier do?")
	pc.Print(doc)
	pc.Scan(doc)
}
