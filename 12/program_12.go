package main

import "fmt"

// removeAtIndex is ..
func removeAtIndex(source []string, index int) []string {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func main() {
	students := []string{"ali", "maul", "kamil"}
	students = removeAtIndex(students, 1)
	fmt.Printf("%s", students)
}
