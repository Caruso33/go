//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func main() {
	greet("Ungar")
	fmt.Println(returningMsg("Yes Ungar!!!"))

	first, second := returningTwoNr(1, 5)
	added3 := add3(first, second, returningOneNr(2))
	fmt.Println(returningMsg("Added 3 numbers"), added3)

}

func greet(name string) {
	fmt.Println("Hi", name)
}

func returningMsg(msg string) string {
	return msg
}

func returningOneNr(nr int) int {
	return nr
}

func returningTwoNr(nr1, nr2 int) (int, int) {
	return nr1, nr2
}

func add3(first, second, third int) int {
	return first + second + third
}
