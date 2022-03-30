package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//This code creates random number and place as variable

	//same seed value will result in the same random sequence each run
	// rand.Seed(11)
	//each time set the seed to different value so it will be different
	//use time.Now().UnixNano() to get constant changing number
	rand.Seed(time.Now().UnixNano())

	//rand.Intn() is half open interval [0-n), n not included
	hiddenNumber := int64(rand.Intn(101))
	fmt.Println(hiddenNumber)

	//Insert code here
	var number string
	count := 5
	contFlag := true

	for contFlag {
		fmt.Println("Guess a number or enter 101 to quit: ")
		fmt.Scanln(&number)

		// if there is error means the string input is not a number
		num, err := strconv.ParseInt(number, 10, 0)

		// if there is error, go back to the starting of for loop to ask for input
		if err != nil {
			continue
		} else {
			// if num is 101, quit the program
			if num == 101 {
				contFlag = false
			} else if num == hiddenNumber {
				fmt.Println("You have guessed correctly the hidden number!")
				contFlag = false
			} else {
				count -= 1
				fmt.Printf("You have %d times left to guess the number\n", count)
			}
		}

		if count == 0 {
			fmt.Println("You have used up your chances!")
			contFlag = false
		}

		}

	}

	// fmt.Println("Guess a number or enter 101 to quit: ")
	// fmt.Scanln(&number)

	// num, err := strconv.ParseInt(number, 10, 0)
	// fmt.Println(num, err)

	// if err != nil {
	// 	fmt.Println("Guess a number: ")
	// }

}
