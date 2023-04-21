package main

import "fmt"

func main() {
	mybill := newBill("Mario's bill")

	mybill.addItem("coke", 5)
	mybill.updateTip(10)
	fmt.Println(mybill.format())

}
