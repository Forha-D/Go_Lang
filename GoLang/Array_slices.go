package main

import "fmt"

func main() {

	var age [3]int = [3]int{20, 25, 30}
	//[3]int = [3]int{20,25,30} same as upper line
	var name [4]string = [4]string{"mario", "peach", "yoshi", "bowser"}
	name[2] = "selina"

	fmt.Println(age, len(age))
	fmt.Println(name, len(name))

	//slices (use under the hood)

	scores := []int{100, 200, 300}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//slice ranges

	rangeOne := name[1:3]
	rangeTwo := name[2:]
	rangeThree := name[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "kappa", "murph", " maria", "yohi")
	fmt.Println(rangeOne)
}
