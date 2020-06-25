package main

import "fmt"

func main() {
	scores := [4]int{1000, 2000, 3000, 4000}
	for _, value := range scores {
		fmt.Printf("%d\n", value)
	}
	warriors := make([]string, 0, 5)
	warriors = append(warriors, "Gimli")
	fmt.Printf("%s", warriors)
	students := make([]string, 0, 5)
	students = students[0:3]
	students[1] = "Ali"
	fmt.Printf("%s", students)
}
