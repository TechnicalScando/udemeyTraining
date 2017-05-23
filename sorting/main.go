package main

import (
	"fmt"
	"sort"
)

func main() {

	type people []string
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println("Unsorted Study Group:", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println("Sorted Study Group:", studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}

	fmt.Println("Unsorted slice of string: ", s)
	sort.Strings(s)
	fmt.Println("Sorted slice of string: ", s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println("Unsorted slice of ints: ", n)
	sort.Ints(n)
	fmt.Println("Sorted slice of int: ", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("Reversed slice of int: ", n)
}
