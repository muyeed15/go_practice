package scopes

import "fmt"

// package scope
var z = 4

func Run() {
	// block scope
	{
		x := 2
		fmt.Println(x)
	}

	y := 3
	fmt.Println(y)

	fmt.Println(z)

}

// scope
// block --> package --> universe
