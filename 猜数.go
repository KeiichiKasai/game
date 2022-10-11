package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
FLAG:
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("Please input your guess")
	for {
		var guess int
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Invalid input.Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number.")
			var G string
			fmt.Println("Give up？ Y/N")
			fmt.Scanln(&G)
			if G == "Y" {
				fmt.Println("The secretNumber is ", secretNumber)
				break
			} else if G == "N" {
				fmt.Println("Please input your guess")
				continue

			}
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number.")
			var G string
			fmt.Println("Give up？ Y/N")
			fmt.Scanln(&G)
			if G == "Y" {
				fmt.Println("The secretNumber is ", secretNumber)
				break
			} else if G == "N" {
				fmt.Println("Please input your guess")
				continue
			}
		} else {
			fmt.Println("Correct,you Legend!")
			break
		}

	}
	fmt.Println("Do you want to play again?  Y/N")
	var A string
	fmt.Scanf("%s\n", &A)
	fmt.Println()
	if A == "Y" {
		goto FLAG
	} else {
		fmt.Println("Bye~")
	}

}
