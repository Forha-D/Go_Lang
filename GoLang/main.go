package main

import "fmt"

func main() {

	age := 35
	name := "mario"

	fmt.Println("my age is ", age, "My name is ", name, "\n")
	fmt.Printf("My age is %v My name is %v \n", age, name)
	var str = fmt.Sprintf("My age is %v My name is %v \n", age, name)
	fmt.Println(str)

}
