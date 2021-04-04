package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

// function for find missing number
func findLostNumber(input string, kali int) (int, bool) {
	var idxStartLast int
	var idxNextfirst int
	var idxNextLast int
	var pindahDigit bool = false
	moveCount := 0
	var i int
	result := 0
	err := true


	for i=0; i<=len(input);i+=kali {	
	
			if pindahDigit == false {
				// second index of first number
				if i + kali < len(input) {
					idxStartLast = i + kali
				} else {
					idxStartLast = len(input)
				}

				// first index of second number

				if i + kali < len(input) {	
					idxNextfirst = i + kali
				} else {
					idxNextfirst = len(input)
				}
		
				// second index of second number
				if idxNextfirst + kali < len(input) {
					idxNextLast = idxNextfirst + kali	
				} else {
					idxNextLast = len(input)
				}
			} else {
				moveCount += 1

				if moveCount > 1 {
					i+=1
				}


				// second index of first number
				if i + kali < len(input) {
					idxStartLast += 2
				} else {
					idxStartLast = len(input)
				}

				// first index of second number

				if i + kali < len(input) {	
					idxNextfirst += 2
				} else {
					idxNextfirst = len(input)
				}
		
				// second index of second number
				if idxNextfirst + kali < len(input) {
					idxNextLast += 2
				} else {
					idxNextLast = len(input)
				}
			}

			// add number digit
			if input[i:idxStartLast] == "9" || 
			input[i:idxStartLast] == "99" || 
			input[i:idxStartLast] == "999" || 
			input[i:idxStartLast] == "9999" || 
			input[i:idxStartLast] == "99999" || 
			input[i:idxStartLast] == "999999"  {
				idxNextLast += 1
				pindahDigit = true
			}

		// convert input to integer
		firstNum, _ := strconv.Atoi(input[i:idxStartLast])
		secNum, _ := strconv.Atoi(input[idxNextfirst:idxNextLast])

		// is number in sequence ?
		if secNum == firstNum + 1 {
			continue
		} else if secNum == firstNum + 2 {
			result = firstNum + 1
			return result, false
		} else {
			break
		}

	}
		return result, err
}

func loop(input string) (int, bool) {
	var result int
	var err bool = true

	// looping number for run finLostNumber function
	for i:=1; i<=6; i++ {
		result, err = findLostNumber(input, i)
		if err == false {
			return result, err
		}
	}

	return result, err
}

func main() {
	fmt.Println("Welcome to find number app")
	fmt.Println("enter a list of number without space that you want to search, number must be in sequence. ex : 1112131516171819")
	fmt.Println("Make sure the list of number that you insert is in correct format. If not program will automatically close :)")

	// scan input from terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\ninput : ")
	scanner.Scan()
	input := scanner.Text()

	// run loop function with input as a param
	result, err := loop(input)

	// if no error print result
	if err == false {
		fmt.Println("result :", result)
	} else {
		fmt.Println("result : not found :(")
	}
}