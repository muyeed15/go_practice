package forloops

import "fmt"

func Run() {
	// for loop style 1
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	// for loop style 2
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// index print
	names := []string{"Muyeed", "Dipra", "Alam"}

	for j := 0; j < len(names); j++ {
		fmt.Println(names[j])
	}

	// for in with index and value
	for index, value := range names {
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

	// for in with only with value
	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
	}
}
