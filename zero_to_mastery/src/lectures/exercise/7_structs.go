//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	a, b, c, d int
}

func calcArea(rect Rectangle) int {
	return (rect.b - rect.a) * (rect.d - rect.c)
}

func calcPerimeter(rect Rectangle) int {
	return (2*(rect.b-rect.a) + 2*(rect.d-rect.c))
}

func printInfo(rect Rectangle) {
	fmt.Println("Area is", calcArea(rect))
	fmt.Println("Perimeter is", calcPerimeter(rect))

}

func main() {
	rect := Rectangle{6, 8, 3, 4}
	// rect := Rectangle{a: 6, b: 8, c: 3, d: 4}
	printInfo((rect))
}
