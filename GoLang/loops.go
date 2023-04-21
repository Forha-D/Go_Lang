package main

import (
	"fmt"
	"sort"
)

func main() {

	//x := 0

	//for x < 5 {

	//fmt.Println("value of the x is :", x)
	//x++
	//}

	//for i := 0; i < 5; i++ {
	//fmt.Println("value of the i is :", i)
	//}

	name := []string{"mario", "peach", "yoshi", "bowser"}
	//for i := 0; i < len(name); i++ {
	//sort.Strings(name)
	//fmt.Println(name[i])
	//}

	//for index, value := range name {
	//	sort.Strings(name)
	//	fmt.Printf("the value of index %v is %v \n", index, value)

	//}
	for _, value := range name {
		sort.Strings(name)
		fmt.Printf("the value is %v \n", value)

	}

}
