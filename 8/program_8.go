package main

import "fmt"

// Saiyan is..
type Saiyan struct {
	Name  string
	Power int
}

// NewSaiyan is..
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

// String is..
func (saiyan *Saiyan) String() string {
	return fmt.Sprintf("<%p Name=%s Power=%d>", &saiyan, saiyan.Name, saiyan.Power)
}

func main() {
	goku := NewSaiyan("Goku", 9000)
	fmt.Printf("%s", goku)
}
