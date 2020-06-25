package main

import (
	"fmt"
)

func main() {
	power := 9000
	fmt.Printf("%d\n", power)
	power = addPower(power, 1000)
	fmt.Printf("%d\n", power)
	power, _ = fetchPower("Goku")
	fmt.Printf("%d\n", power)
}

func addPower(a, b int) int {
	return a + b
}

func fetchPower(name string) (int, bool) {
	return 9000, false
}
