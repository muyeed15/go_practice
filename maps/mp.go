package maps

import "fmt"

func Run() {

	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"tofee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		1712345678: "Muyeed",
		1812345678: "Dipra",
		1912345678: "Alam",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1812345678])

	phonebook[1912345678] = "Romeo"
	fmt.Println(phonebook)

}
