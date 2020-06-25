package main

import (
	"fmt"
)

// Saiyan is
type Saiyan struct {
	Name  string
	Power int
}

// Super is
func (saiyan *Saiyan) Super() {
	saiyan.Power += 10000
}

func main() {
	goku := &Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	goku.Super()
	fmt.Printf("%d", goku.Power)
}
