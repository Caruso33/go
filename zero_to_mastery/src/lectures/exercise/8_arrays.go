//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func getStats(list []Product) (float32, int) {
	var cost float32
	var totalItems int

	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			totalItems++
		}
	}

	return cost, totalItems
}

func main() {
	shopping_list := [...]Product{
		{name: "Shampoo", price: 2.99},
		{name: "Cucumber", price: 1.49},
		{name: "Bottle of Wine", price: 12.49},
	}

	cost, totalItems := getStats(shopping_list[:])

	fmt.Println("Last item", shopping_list[len(shopping_list)-1])
	fmt.Println("Number of items", totalItems)
	fmt.Println("Total cost of list", cost)
}
