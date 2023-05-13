package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var (
	keys      [2]rune
	debugMode = true
)

func main() {
	getKeys()
	debugPrint(fmt.Sprintf("Keys: %v %v", string(keys[0]), string(keys[1])))

	// get keyboard inputs
	err := keyboard.Open()
	panicErr(err)
	defer keyboard.Close()

	// prompt to tell user what to do
	fmt.Println("Please begin alternating, press <Esc> to quit at any time")

	for {
		char, key, err := keyboard.GetSingleKey()
		panicErr(err)

		// print the current key
		debugPrint(fmt.Sprintf("You pressed: %v, %v", string(char), key))

		if key == keyboard.KeyEsc {
			debugPrint("Quitting...")
			break
		}
	}
}

func getKeys() {
	// this function will get the user to enter their 2 desired keys to alternate
	var input string
	for {
		for i := 0; i < 2; i++ {
			for {
				fmt.Printf("Enter key for key %v\n", i+1)
				fmt.Scanln(&input)

				// make sure that user only entered 1 key
				if len(input) == 1 {
					// input[0] returns a byte of the character and then that's converted to a rune
					keys[i] = rune(input[0])
					break
				}
				// this means that the input was not valid
				fmt.Println("Invalid input")
			}
		}
		if keys[0] != keys[1] {
			break
		}
		// tell user that they cannot have the same key twice
		fmt.Println("You can't have the same key twice!")
	}
}

func debugPrint(printStatement string) {
	if debugMode {
		fmt.Println("[DEBUG]", printStatement)
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func isKey(pressed string) {
}
