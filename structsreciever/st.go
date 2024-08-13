package structsreciever

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

// reciever function
func (b bill) getName() string {

	return b.name

}

func (b bill) format() string {

	fs := fmt.Sprintf("Biller Name%-17v : %v\n\nBilled Items", "", b.name)

	sum := 0.0
	i := 0

	for prodcut, cost := range b.items {
		i += 1
		sum += cost
		fs += fmt.Sprintf("\n%v) %-25v : %vTk", i, prodcut, cost)
	}

	fs += fmt.Sprintf("\n\nTip%-25v : %vTk\n\nTotal bill%-18v : %0.2fTk", "", b.tip, "", sum+b.tip)

	return fs

}

// pointer for update
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(product string, cost float64) {
	b.items[product] = cost
}

func Run() {

	billInfo := newBill(
		"Muyeed",
		map[string]float64{"Coconut": 100,
			"Orange": 300,
		},
		20.54,
	)

	fmt.Println(billInfo)
	fmt.Println(billInfo.getName())
	billInfo.updateTip(15.67)
	billInfo.addItem("Milk", 110)
	billInfo.addItem("Corn", 40)
	fmt.Println(billInfo.format())

}
