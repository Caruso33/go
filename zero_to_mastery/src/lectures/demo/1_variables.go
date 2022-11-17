package main

import "fmt"

func main() {
	var myName = "Tobias"
	fmt.Println("my name is", myName)

	var name string = "Kathy"
	userName := "admin"
	fmt.Println("name is", name, userName)

	var sum int
	fmt.Println("sum is", sum)

	part1, other := 1, 5
	fmt.Println("part is", part1, "other is", other)

	part2, other := 2, 6
	fmt.Println("part is", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("sum is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lesson is", lessonName, "type is", lessonType)

	word1, word2, _ := "Hello", "world", "!"
	fmt.Println(word1, word2)
}
