package functions

import (
	"fmt"
	"math"
	"strings"
)

func sayName(name string) {
	fmt.Printf("My name is %v\n", name)
}

func sayAge(age int) {
	fmt.Printf("I am %v\n", age)
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// calling a function for each element in array
func cycleNames(names []string, function func(string)) {
	for _, value := range names {
		function(value)
	}
}

// multi value return
func getMultiValue(word string) (string, string) {
	upper := strings.ToUpper(word)
	names := strings.Split(word, " ")

	return upper, names[0]
}

func Run() {
	sayName("Muyeed")
	sayAge(23)
	cycleNames([]string{"Muyeed", "Dipra", "Alam"}, sayName)

	area := circleArea(10.5)
	fmt.Println(area)
	fmt.Printf("%0.2f\n", area)

	value1, value2 := getMultiValue("Al Muyeed")
	fmt.Println(value1, value2)
}
