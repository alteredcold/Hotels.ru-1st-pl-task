package main

import (
	"fmt"
	"strconv"
)

func getComputerPhrase(count int) string {
	lastNumber := count % 10
	lastTwoNumbers := count % 100

	var word string
	switch {
	case lastTwoNumbers >= 11 && lastTwoNumbers <= 14:
		word = "компьютеров"
	case lastNumber == 1:
		word = "компьютер"
	case lastNumber >= 2 && lastNumber <= 4:
		word = "компьютера"
	default:
		word = "компьютеров"
	}

	return "Go: " + strconv.Itoa(count) + " " + word
}

func main() {
	fmt.Println(getComputerPhrase(25))
	fmt.Println(getComputerPhrase(41))
	fmt.Println(getComputerPhrase(1048))
}
