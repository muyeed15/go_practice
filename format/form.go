package format

import "fmt"

func Run() {
	var name = "Muyeed"
	var age = 23

	// print
	fmt.Printf("My name is %v and I'm %v years old.", name, age)

	// type
	fmt.Printf("\n%T", age)

	// float
	fmt.Printf("\n%f", 12.44565)
	fmt.Printf("\n%0.1f", 12.44565)
	fmt.Printf("\n%0.2f", 12.44565)
	fmt.Printf("\n%0.3f", 12.44565)

	// sprintf
	var str = fmt.Sprintf("My name is %v and I'm %v years old.", name, age)
	fmt.Println("\nSaved String: ", str)
}