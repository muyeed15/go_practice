package main

import (
	"fmt"
	"math"
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

func functions() {
	sayName("Muyeed")
	sayAge(23)
	cycleNames([]string{"Muyeed", "Dipra", "Alam"}, sayName)
	
	area := circleArea(10.5)
	fmt.Println(area)
	fmt.Printf("%0.2f", area)
}
