// comment jajaja
package main

/*
TODO:
[ ] Make it run in the background
*/

import (
	"fmt"
	"os"
	"strconv"
	// "strings"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

var (
	keys        [2]rune
	previousKey rune
	debugMode   bool
	flags       = make([]string, 0, 10)

	// toleration
	tolarance int
	tolarated int

	// counter
	streak int

	// colors
	colors        = [...]*color.Color{color.New(color.FgHiBlue), color.New(color.FgHiMagenta)}
	negativeColor = color.New(color.FgHiRed)
	positiveColor = color.New(color.FgHiGreen)
	exceedColor   = color.New(color.FgBlack).Add(color.BgHiRed)
	warningColor  = color.New(color.FgBlack).Add(color.BgYellow)
	debugColor    = color.New(color.FgBlack).Add(color.BgGreen)
	noBgColor     = color.New(color.FgWhite)
)

func main() {
	getFlags()

	// color for when user gets out of tolaranced threshold

	setToleration()

	// tolerated is automatically set to 0 in go

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

		// in this commit the Escape key won't be triggered as an "undefined key" by checking if the key being pressed is the escape key or not first. This also prevents unintentional streak increases.

		if key == keyboard.KeyEsc {
			debugPrint("Quitting...")
			break
		}

		if char != keys[0] && char != keys[1] {
			warningColor.Print("Please only press the defined keys.")
			fmt.Printf("\n")
			streak = 0
			fmt.Println("Streak has been reset")
			tolarated = 0
		} else {
			// check if currently pressed key is the same as the previously pressed key
			if previousKey != 0 && previousKey == char {
				tolarated++
				if tolarated > tolarance {
					streak = 0
					exceedColor.Print("Maximum tolaration reached, streak has been reset.")
					// new line needs to printed like so to avoid the background not being reset
					fmt.Printf("\n")

				} else {
					streak++
				}
			} else {
				tolarated = 0
				streak++

			}

			printStreak()

			// used to compare and make sure that the next key is not the same as previous
			previousKey = char
		}

		debugPrint(fmt.Sprintf("Tolarance is %v and tolarated is %v", tolarance, tolarated))

	}
}

func getKeys() {
	// this function will get the user to enter their 2 desired keys to alternate
	var input string
	for {
		for i := 0; i < 2; i++ {
			for {
				fmt.Printf("Enter key for key %v\n", i+1)
				colors[i].Print("❯ ")
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
		debugColor.Print("[DEBUG]")
		fmt.Println("", printStatement)
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
		fmt.Print("Enter tolarance value (how many repeated keystrokes are allowed)\n")
		positiveColor.Print("❯ ")
		fmt.Scanln(&input)

		// convert input into an integer
		intInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid integer")
		} else {
			tolarance = intInput
			break
		}
	}
}

func printStreak() {
	fmt.Print("Streak: ")

	if streak > 0 {
		positiveColor.Print(streak)
	} else {
		negativeColor.Print(streak)
	}
	fmt.Print("\n")
}

func getFlags() {
	// This function will check whether -v or --verbose flag has been passed through and if so turn on debug mode

	arguments := os.Args

	for _, value := range arguments {
		// fmt.Println(index, value)
		intString := string(value[0])

		// fmt.Println("This is the first character:", intString)
		if intString == "-" {
			// fmt.Println("This string has a - infront of it:", value)
			flags = append(flags, value)
		}
	}

	setDebug()

	debugPrint(fmt.Sprintf("Length of args: %v\nArguments: %v\n", len(os.Args), arguments))
}

func setDebug() {
	for _, value := range flags {
		if value == "-v" || value == "--verbose" {
			debugMode = true
		}
	}
}
