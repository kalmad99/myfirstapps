package main

import (
	"fmt"
	"strconv"
)

func main() {
	var rows int = 12
	var result int

	for i := 1; i <= rows; i++ {
		fmt.Println("Multiples of " + strconv.Itoa(i) + " are: ")
		for j := 1; j <= rows; j++ {
			result = i * j
			fmt.Print(strconv.Itoa(result) + " ")
		}
		fmt.Println(" ")
	}
}

// package main

// import "fmt"

// func main() {
// 	var n int

// 	fmt.Println("Type '0' if you want to end the game!!")

// 	for {

// 		fmt.Print("Enter any non-zero integer: ")

// 		fmt.Scan(&n)

// 		if n == 0 {
// 			fmt.Println("Thank You")
// 			break
// 		} else {
// 			for i := 1; i <= 12; i++ {
// 				fmt.Println(n, "x", i, " = ", n*i)
// 			}
// 		}
// 	}
// }
