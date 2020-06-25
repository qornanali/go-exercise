package main

import "fmt"

// Person is ..
type Person struct {
	Name string
}

// Introduce is ..
func (person *Person) Introduce() string {
	return fmt.Sprintf("%s", person.Name)
}

// Saiyan is ..
type Saiyan struct {
	*Person
	Power int
}

// String is..
func (saiyan *Saiyan) String() string {
	return fmt.Sprintf("<%p Name=%s Power=%d>", &saiyan, saiyan.Name, saiyan.Power)
}

// NewSaiyan is ..
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Person: &Person{name},
		Power:  power,
	}
}

// Introduce is ..
func (saiyan *Saiyan) Introduce() string {
	return fmt.Sprintf("Hello, my name is %s", saiyan.Name)
}

func main() {
	goku := NewSaiyan("goku", 9000)
	fmt.Printf("%s", goku)
	fmt.Printf("%s", goku.Introduce())
}
