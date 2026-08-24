package main

import (
	"fmt"
	"math/rand"
)

func gameEngine(choice int) {

	correctNumber := rand.Intn(100) + 1

	var (
		attemps   int
		userinput int
	)

	if choice == 1 {
		attemps = 10
	} else if choice == 2 {
		attemps = 7
	} else if choice == 3 {
		attemps = 5
	}

	for i := 1; i <= attemps; i++ {
		fmt.Print("Guess The Number (1-100) : ")
		fmt.Scan(&userinput)

		if userinput == correctNumber {
			fmt.Println("True, Correct Number : ", correctNumber)
			break
		} else if userinput > correctNumber {
			fmt.Println("Wrong, < ", userinput)
		} else if userinput < correctNumber {
			fmt.Println("Wrong, > ", userinput)
		}

		fmt.Println("Attemps : ", i)
		if attemps == i {
			fmt.Print("Game Over")
		}

	}
}

func main() {
	var (
		choice   int
		tryAgain string
	)

	for {
		fmt.Println("Welcome to Guess The Number Game\nSelect Level : \n1. Easy\n2. Medium, \n3. Hard")

		fmt.Print("\n Level Choice : ")
		fmt.Scan(&choice)
		if choice >= 1 && choice <= 3 {
			gameEngine(choice)
		} else {
			fmt.Println("Wrong Choice")
		}

		fmt.Print("\nTry again? y/n : ")
		fmt.Scan(&tryAgain)
		if tryAgain == "y" || tryAgain == "Y" {
			continue
		} else if tryAgain == "n" || tryAgain == "N" {
			break
		} else {
			fmt.Println("WRONG ANSWER")
		}

	}
}
