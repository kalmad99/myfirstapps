package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	userInput string
	temp      int64
	attempts  int   = 5
	constant  int64 = 0
)

func numRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	numberToGuess := numRandom(1, 20)
	// fmt.Println(numberToGuess)

	for attempts >= 1 {
		attempts--
		fmt.Println("A number is generated between 1 to 20. What do you think the number is: ")
		fmt.Scanf("%v\n", &userInput)

		// check if userinput is integer
		userNumber, err := strconv.ParseInt(userInput, 10, 0)
		constant = userNumber

		if err != nil {
			fmt.Println("You must enter an integer value.")
			attempts++
		} else if numberToGuess == int(userNumber) {
			fmt.Println("Correct")
			break
		} else if constant == temp {
			attempts++
			fmt.Printf("You are guessing the same number. You have %d attempts left\n", attempts)
		} else if int(userNumber) < numberToGuess {
			fmt.Printf("Smaller than the number. You have %d attempts left\n", attempts)
		} else {
			fmt.Printf("Larger than the number. You have %d attempts left\n", attempts)
		}
		temp = constant
	}
	if int(constant) == numberToGuess {
		fmt.Printf("You win. The correct number was %d\n", numberToGuess)
		fmt.Printf("You had %d guesses left. Good job\n", attempts)
	} else {
		fmt.Printf("You lose! The right answer was : %d\n", numberToGuess)
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	guessLeft := 5
// 	var guess int
// 	secretNumber := 12
// 	fmt.Println("I am thinking of a number from 0 to 20. You will have 5 tries!")
// 	fmt.Println("Good Luck!")
// 	for guessLeft > 0 {
// 		fmt.Scanf("%d", &guess)
// 		guessLeft--

// 		if guess == secretNumber {
// 			break
// 		} else if guess < secretNumber {
// 			fmt.Println("Go higher")
// 			fmt.Printf("You have %d tries left\n", guessLeft)
// 		} else if guess > secretNumber {
// 			fmt.Println("Go Lower")
// 			fmt.Printf("You have %d tries left\n", guessLeft)
// 		}
// 	}
// 	if guess == secretNumber {
// 		fmt.Printf("You win. The correct number was %d\n", guess)
// 		fmt.Printf("You had %d guesses left. Good job\n", guessLeft)
// 	} else {
// 		fmt.Printf("You lose! The right answer was : %d\n", secretNumber)
// 	}
// }
