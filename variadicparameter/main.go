package main

import "fmt"

func main() {
	fmt.Println(greatestNumber(17, 2, 3, 4, 5, 6, 9, 15))
}

func greatestNumber(n ...int) int {
	var max int
	for _, v := range n {
		if max < v {
			max = v
		}
	}
	return max
}
