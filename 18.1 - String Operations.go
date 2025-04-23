/*
 * Author      : Luke Sathrum
 * Description : Examples of string operators
 */

package main

import "fmt"

func main() {
	// Use = for assignment
	var str = "Hello"
	var str2 = "World!"

	// Use + for concatenation
	str3 := str + " " + str2

	// Output the concatenated string
	fmt.Println(str3)

	// Same string but lowercased
	strLower := "hello world!"

	// Compare the String (Lexiographic Ordering)
	if str3 < strLower {
		fmt.Println(str3 + " is smaller")
	} else {
		fmt.Println(strLower + " is smaller")
	}

	// Strings start out empty
	var empty string

	// Test for the "Empty String"
	if empty == "" {
		fmt.Println("There are no characters in empty")
	}
}
