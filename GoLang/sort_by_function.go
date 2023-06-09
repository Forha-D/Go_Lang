package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
func main() {
	fruits := []string{"banana", "peach", "kiwi"}
	sort.Strings(fruits)
	fmt.Println(fruits)
}
