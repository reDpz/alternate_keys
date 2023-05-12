package main

import (
	"fmt"
)

var (
	keys      [2]string
	debugMode = true
)

func main() {
	getKeys()
	fmt.Printf("Keys: %v %v", keys[0], keys[1])
}

func getKeys() {
	var input string
	for i := 0; i < 2; i++ {
		for {
			fmt.Printf("Enter key for key %v\n", i)
			fmt.Scanln(input)

			// make sure that user only entered 1 key
			if len(input) == 1 {
				keys[i] = input
				break
			}
			// this means that the input was not valid
			fmt.Println("Invalid input")
		}
	}
}

func getKeys
