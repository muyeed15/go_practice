package variablesstringsnumbers

import "fmt"

func RunVar() {

	// strings
	var nameOne string = "Muyeed"
	var nameTwo = "Dipra"

	// use for future
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// updating variable
	nameOne = "Alam"
	nameThree = "Romeo"

	fmt.Println(nameOne, nameThree)

	// without var keyword
	nameFour := "Diya"

	fmt.Println((nameFour))

}
