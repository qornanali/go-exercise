package main

import "fmt"

// Saiyan is..
type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

// NewSaiyan is..
func NewSaiyan(name string, power int, father *Saiyan) *Saiyan {
	return &Saiyan{
		Name:   name,
		Power:  power,
		Father: father,
	}
}

// String is..
func (saiyan *Saiyan) String() string {
	return fmt.Sprintf("<%p Name=%s Power=%d>", &saiyan, saiyan.Name, saiyan.Power)
}

func main() {
	goku := NewSaiyan("Goku", 9000, nil)
	gohan := NewSaiyan("Gohan", 2000, goku)
	fmt.Printf("%s", goku)
	fmt.Printf("%s", gohan)
}
