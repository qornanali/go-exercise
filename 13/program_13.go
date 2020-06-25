package main

import "fmt"

func main() {
	dictionary := make(map[string]int)
	dictionary["goku"] = 9000
	power, exist := dictionary["goku"]

	fmt.Println(power, exist)

	total := len(dictionary)

	fmt.Println(total)

	delete(dictionary, "goku")

	power, exist = dictionary["goku"]
	fmt.Println(power, exist)
}
