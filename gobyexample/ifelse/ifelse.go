package main

import "fmt"

func main() {
	// if else is mostly the same,
	// this is the only "weird" if/else combo

	// basically go's version of a ternary conditional
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
