package switchcase

import "fmt"

func selectOption(ui string) {
	switch ui {
	case "a":
		fmt.Println("You chose a")
	case "t":
		fmt.Println("You chose t")
	default:
		fmt.Println("You should chose a or t")
	}
}

func Run(ui string) {
	selectOption(ui)
}
