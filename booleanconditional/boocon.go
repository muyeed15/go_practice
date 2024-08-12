package booleanconditional

import "fmt"

func Run() {
	age := 45

	// direct bool
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	// condtion
	if age < 30 {
		fmt.Println("true")
	} else if age == 45 {
		fmt.Println("yuppies!")
	} else {
		fmt.Println("false")
	}

	// continue and break
	names := []string{"Muyeed", "Dipra", "Alam", "Romeo", "Diya", "Shrabon"}

	for index, value := range names {
		if index == 2 {
			fmt.Println("Continuing at pos", index)
			continue
		}

		if index == 4 {
			fmt.Println("Breaking at pos", index)
			break
		}

		fmt.Printf("The value of pos %v is %v\n", index, value)

	}

}
