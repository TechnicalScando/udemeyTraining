package main

import "fmt"

func main() {
	fmt.Println(doublereturn(1))
	fmt.Println(doublereturn(2))
}

func doublereturn(n int) (int, bool) {
	var d int
	var b bool

	d = n / 2

	if n%2 == 0 {
		b = true
	}

	return d, b
}
