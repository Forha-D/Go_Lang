package main

import "fmt"

func option(b bill) {

	opt := " choose option (a	- add item, s	- save bill , t	- add tip )"
	fmt.Println(opt)
	var opt1 string
	fmt.Scanln(&opt1)

	switch opt1 {

	case "a":
		item := "item name"
		fmt.Println(item)
		var n string
		fmt.Scanln(&n)
		price := "item price"
		fmt.Println(price)
		var ip float64
		fmt.Scanln(&ip)
		fmt.Println(n, ip)

	case "s":
		b.save()
		fmt.Println("you save the bill", b.name)
	case "t":
		tip1 := "Enter tip amount of ($)	:	"
		fmt.Print(tip1)
		var tip float64
		fmt.Scan(&tip)
		fmt.Println(tip)

	default:
		fmt.Println("Thar was not valid")
		option(b)
	}

}

func main() {

	mybill := newBill(" mario's bill")

	option(mybill)

	//fmt.Println(mybill.format())
}
