package main

import (
	"errorhandling"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	n, err := strconv.Atoi(os.Args[1])
	err = errorhandling.InputError{
		Message: "input is error",
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
