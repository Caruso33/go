//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rolls, sides := 10, 6

	var total int = 0

	for i := 1; i <= rolls; i++ {
		dice := rollDice(sides)

		total += dice

		if dice == 2 && total == 2 {
			fmt.Println("Snake eyes")
		}

		if total == 7 {
			fmt.Println("Lucky 7")
		}

		if total%2 == 0 {
			fmt.Println("Even")
		}

		if total%2 == 1 {
			fmt.Println("Odd")
		}

		fmt.Println("Roll #", i, "Sides #", sides, "Dice", dice, "Total", total)
	}
}

func rollDice(sides int) int {
	return rand.Intn(sides) + 1
}
