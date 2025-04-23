/*
 * Author      : Luke Sathrum
 * Description : Reading and writing files while checking to make sure the files
 *               are opened. Depends on input.txt
 */
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Try to read a file in the current directory
	fileData, err := ioutil.ReadFile("input.txt")
	// Check to make sure the file opened successfully
	if err == nil {
		// The file was opened
		fmt.Printf("%T\n", fileData)
		fmt.Println(string(fileData))
	} else {
		// The file ws NOT opened
		fmt.Println("Could not open file")
	}

	// Write OR Overwrite a file
	var str string
	fmt.Println("Enter a word to write to a file")
	fmt.Scan(&str)
	// Write to the file
	err = ioutil.WriteFile("tmp.txt", []byte(str), 0644)
	// Check to make sure the write worked
	if err == nil {
		// The file was written to
		fmt.Println("Write Successful")
	} else {
		// The file was NOT written to
		fmt.Println("Write Unsuccessful")
	}

	// Finish our rhyme
	rhyme := []byte("\n9 10 begin again.")
	err = ioutil.WriteFile("rhyme.txt", append(fileData, rhyme...), 0644)
	if err == nil {
		// The file was written to
		fmt.Println("Rhyme Write Successful")
	} else {
		// The file was NOT written to
		fmt.Println("Rhyme Write Unsuccessful")
	}

}
