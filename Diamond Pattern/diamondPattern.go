package main

import (
	"fmt"
)

func main() {
	var rows int = 10

	for i := 1; i <= rows; i++ {
		for j := rows; j >= i; j-- {
			fmt.Print(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i := rows; i >= 1; i-- {
		for j := rows + 1; j >= i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k < i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
