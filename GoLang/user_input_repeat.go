package main

import (
	"bufio"
	"fmt"
	"os"
)

func userInput() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s\n", name)

}

func main() {
	userInput()
}
