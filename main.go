// comment jajaja
package main

import (
	"fmt"
	"strconv"

	"github.com/eiannone/keyboard"
)

var (
	keys        [2]rune
	previousKey rune
	debugMode   = true
	tolarance   int
	tolarated   int
)

func main() {
	tolarance = 2
	tolarated = tolarance
	getKeys()
	debugPrint(fmt.Sprintf("Keys: %v %v", string(keys[0]), string(keys[1])))
	debugPrint(fmt.Sprintf("Tolarance is %v and tolarated is %v", tolarance, tolarated))

	// get keyboard inputs
	err := keyboard.Open()
	panicErr(err)
	defer keyboard.Close()

	// prompt to tell user what to do
	fmt.Println("Please begin alternating, press <Esc> to quit at any time")

	// var previousKey rune

	// this is where we actually keep track of key presses
	for {
		debugPrint(fmt.Sprintf("Previous key is %v", previousKey))

		char, key, err := keyboard.GetSingleKey()
		panicErr(err)

		// print the current key
		debugPrint(fmt.Sprintf("You pressed: %v, %v", string(char), key))

		// check if key is in the keys arrray
		if char != keys[0] && char != keys[1] {
			fmt.Println("Please only press the defined keys.")
		} else {
			// check if currently pressed key is the same as the previously pressed key
			if previousKey != 0 && previousKey == char {
				tolarated--
				if tolarated < tolarance {
					fmt.Println("Maximum tolaration reached.")
				}
			} else {
				tolarated = tolarance
			}

			// used to compare and make sure that the next key is not the same as previous
			previousKey = char
		}

		debugPrint(fmt.Sprintf("Tolarance is %v and tolarated is %v", tolarance, tolarated))

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

func setToleration() {
	var input string
	for {
		fmt.Println("Enter tolerance value (how many repeated keystrokes are allowed)")
		fmt.Scanln(&input)

		// convert input into an integer
		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid integer")
		} else {
			tolarance = intInput
		}
	}
}
