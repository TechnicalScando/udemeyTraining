package main

import "fmt"

func main() {
	var smaller int
	var larger int

	fmt.Print("Please enter a number:")
	fmt.Scan(&smaller)
	fmt.Print("Please enter a larger number:")
	fmt.Scan(&larger)

	fmt.Println(larger % smaller)
}
