package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) bill {

	b := bill{
		name: name,
		item: map[string]float64{"pie": 2.00, "salad": 3.00, "cake": 4.00},
		tip:  0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	//list item
	for k, v := range b.item {
		fs += fmt.Sprintf("%v ....$%v \n", k, v)
		total += v
	}
	// total
	fs += fmt.Sprintf("%v...$%0.2f", "total", total)

	//tip
	fs += fmt.Sprintf("%v...$%0.2f", "tip", total+b.tip)
	return fs
}

// update tip
func (b bill) updateTip(tip float64) {
	b.tip = tip

}
