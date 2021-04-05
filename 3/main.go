package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

// func for referse integer
func reverseInt(value int) int {
	// conver int to string
    intString := strconv.Itoa(value)

    newString := ""

	// loop the string and reverse it by index
    for x := len(intString); x > 0; x-- {
        newString += string(intString[x - 1])
    }

	// conver to integer
    newInt, err := strconv.Atoi(newString)

    if(err != nil){
        fmt.Println("Error converting string to int")
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



func palindrom(w http.ResponseWriter, r *http.Request) {
    input1 := 1
    input2 := 10
    result := loop(input1, input2)

    input1Str := strconv.Itoa(input1)
    input2Str := strconv.Itoa(input2)
    resultStr := strconv.Itoa(result)

    fmt.Fprintf(w, "Palindrom, input : '" + input1Str + " " + input2Str + "' result = " + resultStr) // write data to response
}


func main() {
    http.HandleFunc("/", palindrom) // setting router rule
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}