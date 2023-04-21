package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99

}

func main() {

	name := "mario"
	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64{

		"soup":  3.5,
		" pie ": 4.5,
	}
	updateMenu(menu)
	fmt.Println(menu)

}
