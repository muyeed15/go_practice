package structs

import "fmt"

// struct aka blue print
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// function to return the struct aka blue print
func newBill(name string, items map[string]float64, tip float64) bill {

	b := bill{
		name:  name,
		items: items,
		tip:   tip,
	}

	return b

}

func Run() {
	billInfo := newBill(
		"Muyeed",
		map[string]float64{"Coconut": 100,
			"Orange": 300,
		},
		20.2,
	)

	fmt.Println(billInfo)
}
