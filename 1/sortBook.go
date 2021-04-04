package main

import (
	"fmt"
	"strings"
	// "strconv"
	"sort"
	// "reflect"
)

type Book struct {
	Category, Title, Size string
}

func toStruct(input string) Book {
	var book Book
	book.Category = input[0:1]
	book.Title = input[1:2]
	book.Size = input[2:4]

	return book
}

// func (book Book) collection() string {
// 	six := []Book{}
// 	seven := []Book{}

// 	if book.Category == 6 {
// 		result := append(six, book)
// 		fmt.Println(result)		
// 	} else {
// 		result := append(seven, book)
// 		fmt.Println(result)
// 	}
// 	fmt.Println(six)
// 	return "hello bapak"
// }

func sortingSlice(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Title < books[j].Title {
			return true
		}
		if books[i].Title > books[j].Title {
			return false
		}
		return books[i].Size > books[j].Size
	})
}

func removeTripleBook(books []Book) []Book{
	var result []Book
	var title1 string
	var title2 string
	var countSame int

	for i:=0; i < len(books); i++ {
		title1 = books[i].Title
		result = append(result, books[i])

		for j:=0; j < len(books); j++ {
			title2 = books[j].Title
			
			if title1 == title2 {
				countSame += 1
			}

			if countSame == 2 {
				continue
			}

		}	
		
		countSame = 0
	}

	return result
}

func toStr(books []Book) string{
	var result string

	if len(books) == 0 {
		return ""
	}

	for i:=0; i<len(books); i++ {
		str := []string{books[i].Category,books[i].Title,books[i].Size}
		result += strings.Join(str, "") + " "
	}

	return result
}



func main() {
	input := "3A13 6X19 9Y20 2C18 1N20 6N20 1M21 1F14 9A21 3N14 0E13 5G14 8A23 9E22 3N21 3N20 3N18"

	split := strings.Split(input, " ")

	var books0, books1, books2, books3, books4, books5, books6, books7, books8, books9 []Book
	
	for i:=0; i<len(split); i++ {
		toStruct := toStruct(split[i])

		if toStruct.Category == "0" {
			books0 = append(books0, toStruct)
		} else if toStruct.Category == "1" {
			books1 = append(books1, toStruct)
		} else if toStruct.Category == "2" {
			books2 = append(books2, toStruct)
		} else if toStruct.Category == "3" {
			books3 = append(books3, toStruct)
		} else if toStruct.Category == "4" {
			books4 = append(books4, toStruct)
		} else if toStruct.Category == "5" {
			books5 = append(books5, toStruct)
		} else if toStruct.Category == "6" {
			books6 = append(books6, toStruct)
		} else if toStruct.Category == "7" {
			books7 = append(books7, toStruct)
		} else if toStruct.Category == "8" {
			books8 = append(books8, toStruct)
		} else if toStruct.Category == "9" {
			books9 = append(books9, toStruct)
		}
	}

	sortingSlice(books1)
	sortingSlice(books2)
	sortingSlice(books3)
	sortingSlice(books4)
	sortingSlice(books5)
	sortingSlice(books6)
	sortingSlice(books7)
	sortingSlice(books8)
	sortingSlice(books9)

	strBook0 := toStr(books0)
	strBook1 := toStr(books1)
	strBook2 := toStr(books2)
	strBook3 := toStr(books3)
	strBook4 := toStr(books4)
	strBook5 := toStr(books5)
	strBook6 := toStr(books6)
	strBook7 := toStr(books7)
	strBook8 := toStr(books8)
	strBook9 := toStr(books9)


	removeTripleBook(books3)

	fmt.Println(strBook6 + "" + strBook7 + "" + strBook0 + "" + strBook9 + "" + strBook4 + "" + strBook8 + "" + strBook1 + "" + strBook2 + "" + strBook5 + "" + strBook3)
	
}