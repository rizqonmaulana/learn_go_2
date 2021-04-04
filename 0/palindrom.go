package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func reverseInt(n int) int {
    newInt := 0
    for n > 0 {
        remainder := n % 10
        newInt *= 10
        newInt += remainder 
        n /= 10
    }
    return newInt 
}

//func palindromCheck is to check is the input a palindrom or not
func palindromCheck(i int, reverse int) bool {
	if i == reverse {
	// if i & reverse value is same return true to tell that the input is a palindrom
		return true
	} else {
	// if i & reverse value is not same return false to tell that the input is not a palindrom
		return false
	}
}

// func errorCheck is for check the input value is it on correct format or not
func errorCheck(int1, int2 int) string {
	if int1 > int2 {
		// check is int1 value is higher than int2 ?
		return "int2 value cannot lower than int1 value"
	} else if int1 < 1 {
		// check is int1 value is lower than 1 ?
		return "int value cannot lower than 1"
	} else if int2 > 1000000000 {
		// check is int2 value is higher than 10^9
		return "int value cannot higher than 10^9"
	} else {
		// if there is no error return nil to tell that there is no error
		return ""
	}
}

// func loop is for looping the input
func loop(int1 int, int2 int) int {
	// default value of result is 0, it will increase if there is a palindrom
	result := 0
	for i := int1; i <= int2; i++ {
		// call reverseInt function to reverse the value from loop
		reverse := reverseInt(i)
		
		// call palindromCheck to check is the value from the loop and reverse is same 
		check := palindromCheck(i,reverse)

		// if the value from loop and reverse is same plus 1 to result
		if check {
			result += 1
		}
	}
	return result
}

// func scanInput is to scan the input and make sure the input is correct
func scanInput() {
	fmt.Println("Palindrom check. to end program type 'end'")
	// run scanner on loop, the scanner always run until break
	for {
		// get input
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("input : ")
		scanner.Scan()
		input := scanner.Text()

		// break the loop and stop the programmer if the input is 'end'
		if input == "end" {
			break
		} 

		// split input by space
		separateStr := strings.Split(input, " ")

		/** 
			make sure that input is only 2 after split
			if input != 2 print the error message
		*/
		if len(separateStr) != 2{
			fmt.Println("WRONG INPUT. input must be 2 integer and separate by 1 space. Ex : 10 50")
		} else {
			// transform string to integer
			int1, err := strconv.Atoi(separateStr[0])
			int2, err := strconv.Atoi(separateStr[1])

			// make sure there is no error when change string to integer
			if err != nil {
				fmt.Println("WRONG INPUT. input must be 2 integer and separate by 1 space. Ex : 10 50")
			} else {
				// run errorCheck function to make sure that the input is right
				err := errorCheck(int1, int2)

				/** 
					check is there any error ? 
					if yes print error message
					if no run the loop function
				*/
				if err != "" {	
					fmt.Println(err)
				} else {
					result := loop(int1, int2)
					fmt.Println("output :",result)
				}
			}
		}
	}
}

// main function is the function that will be executed when the file is executed
func main() {
	// run scanInput function
	scanInput()
}