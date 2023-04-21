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
	dat, err := os.ReadFile("tem.txt")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("tem.txt")
	check(err)
	b1 := make([]byte, 2000)
	n1, err := f.Read(b1)
	check(err)
	fmt.Println(n1, "\n", string(b1[:n1]))

}
