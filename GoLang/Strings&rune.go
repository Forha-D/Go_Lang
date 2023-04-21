package main

import (
	"fmt"
	"unicode/utf8"
)

func idvx(data []string) string {

	st := ""

	for i, v := range data {

		fmt.Printf("%d holds :%q\n", i, v)

	}
	return st

}

func main() {

	const s = "Paradise"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {

		fmt.Printf("%x", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, v := range s {
		fmt.Printf("%d value = %q\n", i, v)

	}

	s1 := []string{"jhon", "peter", " tom", "marry"}

	fmt.Println(idvx(s1))

}
