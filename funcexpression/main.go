package main

import "fmt"

func main() {
	doubleReturn := func(n int) (int, bool) {
		var d int
		var b bool

		d = n / 2

		if n%2 == 0 {
			b = true
		}

		return d, b
	}
	fmt.Println(doubleReturn(1))
	fmt.Println(doubleReturn(2))
}
