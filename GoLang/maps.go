package main

import "fmt"

func main() {

	menu := map[string]float64{

		"soup ":        4.65,
		"salad":        1.5,
		"rice":         6.00,
		"soft drinks ": 3.5,
		"beef ":        7.5,
	}
	fmt.Println(menu)
	fmt.Println(menu["rice"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	menu["salad"] = 2.00
	fmt.Println(menu)

	phonebook := map[int]string{
		+8801737356527: "Mobarak",
		+8801737311835: "Masum",
		+8801722905603: "Hena kabir",
		+8801575733120: "Kanan",
	}

	for k, v := range phonebook {
		fmt.Println(k, "-", v)

	}
	fmt.Println(phonebook[01575733120])
}
