package main

import "fmt"

func main() {

	age := 45
	//boolean
	//fmt.Println(age <= 50)
	//fmt.Println(age >= 50)
	//fmt.Println(age == 45)
	//fmt.Println(age != 50)

	// condition

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}
	name := []string{"mario", "peach", "yohi", "bowser", "steve"}

	for index, value := range name {
		if index == 1 {
			fmt.Println("continuing the pos:\n", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos:\n", index)
			break
		}
		fmt.Printf("the value of index %v is %v \n", index, value)

	}

}
