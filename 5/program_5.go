package main

import (
	"fmt"
)

// Saiyan is an object
type Saiyan struct {
	Name  string
	Power int
}

func main() {
	goku := Saiyan{
		Name:  "Goku",
		Power: 9000,
	}
	Super(goku)
	fmt.Printf("%d", goku.Power)
}

// Super is ..
func Super(s Saiyan) {
	s.Power += 10000
}
