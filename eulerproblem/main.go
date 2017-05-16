package main

import "fmt"

func main() {
	fmt.Println(findSmallestDividend(20)) //232792560

}

func findSmallestDividend(n int) int {
	value := 1
	for {
		var sum int
		for i := n; i > 0; i-- {
			sum += value % i
		}
		if sum == 0 {
			return value
		}
		value++
	}
}

//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
