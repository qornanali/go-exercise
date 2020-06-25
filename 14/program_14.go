package main

import "fmt"

// Saiyan is
type Saiyan struct {
	Name    string
	Power   int
	Friends map[string]*Saiyan
}

func main() {
	krillin := Saiyan{
		Name:  "Krillin",
		Power: 3000,
	}
	vegeta := Saiyan{
		Name:  "Vegeta",
		Power: 11000,
	}
	friends := make(map[string]*Saiyan, 2)
	friends["krillin"] = &krillin
	friends["vegeta"] = &vegeta
	goku := Saiyan{
		Name:    "Goku",
		Power:   9000,
		Friends: friends,
	}
	fmt.Println(goku)
}
