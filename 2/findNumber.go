package main

import (
	"fmt"
	"strconv"
)

func findLostNumber(input string, kali int) (int, bool) {
	var idxStartLast int
	var idxNextfirst int
	var idxNextLast int
	var pindahDigit bool = false
	var i int
	result := 0
	err := true

	for i=0; i<=len(input);i+=kali {		
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

			// pindah digit
			if input[i:idxStartLast] == "9" || 
			input[i:idxStartLast] == "99" || 
			input[i:idxStartLast] == "999" || 
			input[i:idxStartLast] == "9999" || 
			input[i:idxStartLast] == "99999" || 
			input[i:idxStartLast] == "999999"  {
				idxNextLast += 1
				pindahDigit = true
			}

		firstNum, _ := strconv.Atoi(input[i:idxStartLast])
		secNum, _ := strconv.Atoi(input[idxNextfirst:idxNextLast])

		fmt.Println("A idx 1 :", i)
		fmt.Println("A idx 2 :", idxStartLast)
		fmt.Println("B idx 1 :", idxNextfirst)
		fmt.Println("B idx 2 :", idxNextLast)
		fmt.Println(firstNum)
		fmt.Println(secNum)

		fmt.Println("####################END########################")
		
		if secNum == firstNum + 1 {
			if pindahDigit == true {
				idxStartLast += 1
				idxNextfirst += 1
				idxNextLast += 1
				// pindahDigit = false
			}
			// continue
		} else if secNum == firstNum + 2 {
			result = firstNum + 1
			return result, false
		} else {
			break
		}

	}
		return result, err
}

func main() {
	input := "78910213"

	var result int
	var err bool

	// looping max value
	for i:=1; i<=6; i++ {
		result, err = findLostNumber(input, i)
		if err == false {
			fmt.Println("angka yang hilang adalah :", result)
			break
		}
	}
}