package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{"pie": 2.00, "cake": 3.99, "coke": 4.5},
		tip:  0,
	}
	return b
}

// Format bill
func (b *bill) format() string {
	fs := " Breaking dowm bill\n"
	var total float64
	//list bill
	for k, v := range b.item {
		fs += fmt.Sprintf("%v......$%v\n", k, v)
		total += v
	}
	//add tip

	fs += fmt.Sprintf("%v...$%0.2f", "tip", b.tip)

	//add total
	fs += fmt.Sprintf("%v......$%0.2f", "\ntoatl", total+b.tip)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item
func (b *bill) addItem(name string, price float64) {
	b.item[name] = price
}

// save file
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bill/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("The bill was saved to the file ")

}
