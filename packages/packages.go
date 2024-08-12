package main

import (
	"fmt"
	"sort"
	"strings"
)

func packages() {
	
	greeting := "Hello world!"

	fmt.Println(greeting)

	// contains
	fmt.Println(strings.Contains(greeting, "Hello"))

	// replace
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	// upper case
	fmt.Println(strings.ToUpper(greeting))

	// lower case
	fmt.Println(strings.ToLower(greeting))

	// inddex of sub-string
	fmt.Println(strings.Index(greeting, "ll"))

	// split
	fmt.Println(strings.Split(greeting, " "))

	// sort
	numbers := []int{2, 6, 1, 3, 0, 9, 4}

	fmt.Println(numbers)

	sort.Ints(numbers)

	fmt.Println(numbers)

	// search
	index := sort.SearchInts(numbers, 6)
	fmt.Println(index)

	// sort string slices
	names := []string{"Muyeed", "Dipra", "Alam"}
	
	sort.Strings(names)

	fmt.Println(names)

}
