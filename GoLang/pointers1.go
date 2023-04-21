package main

import "fmt"

func Zeroval(ival int) {
	ival = 0
}

func Zeroptr(iptr *int) {

	*iptr = 0
}

func main() {
	i := 1

	fmt.Println("initial:", i)

	//Zeroval(i)
	fmt.Println("zeoval:", i)
	Zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("zeoval:", &i)

}
