package main

import (
	"fmt"
	"os"
)

func check(e error) {

	if e != nil {

		panic(e)
	}

}

func main() {
	d1 := []byte("hhhhhhhh")
	err := os.WriteFile("tem.txt", d1, 0644)
	check(err)
	f, err := os.Create("tem1.txt")
	check(err)

	n2, err := f.Write(d1)
	check(err)

	fmt.Println(n2)

}
