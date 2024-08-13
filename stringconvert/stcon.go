package stringconvert

import (
	"fmt"
	"strconv"
)

func Run() {
	var number string = "1234"
	var statement string = "true"

	numberInt, _ := strconv.ParseInt(number, 10, 64)
	numberIntString := strconv.FormatInt(numberInt, 10)

	numberFloat, _ := strconv.ParseFloat(number, 64)
	numberFloatString := strconv.FormatFloat(numberFloat, 'f', -1, 64)

	statementBool, _ := strconv.ParseBool(statement)
	statementBoolString := strconv.FormatBool(statementBool)

	fmt.Printf("%v is %t", numberInt, numberInt)
	fmt.Printf("\n%v is %t", numberIntString, numberIntString)
	fmt.Printf("\n%v is %t", numberFloat, numberFloat)
	fmt.Printf("\n%v is %t", numberFloatString, numberFloatString)
	fmt.Printf("\n%v is %t", statementBool, statementBool)
	fmt.Printf("\n%v is %t", statementBoolString, statementBoolString)
}
