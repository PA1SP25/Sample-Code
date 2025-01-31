/*
 * Author      : Luke Sathrum
 * Description : Demonstrates how to get input from the user
 *               using CinReader
 * Sources     : CinReader (Created by Boyd Trolinger)
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Create a CinReader
  reader := NewCinReader()


  // Create the variables we will be using
  var intVar int
  var floatVar float64
  var charVar rune
  var stringVar string

  // Getting Integers
  fmt.Println("Please enter an integer value ")
  intVar = reader.ReadInteger()
  fmt.Printf("You entered: %v\n\n", intVar)

  // Limiting the integer range
  fmt.Println("Please enter another integer in the range of 1-10 ")
  intVar = reader.ReadIntegerRange(1, 10)
  fmt.Printf("You entered: %v\n\n", intVar)

  // Getting floats
  fmt.Println("Please enter a decimal value ")
  floatVar = reader.ReadFloat()
  fmt.Printf("You entered: %v\n\n", floatVar)


  // Getting runes (Characters)
  fmt.Println("Please enter a single character ")
  charVar = reader.ReadCharacter()
  fmt.Printf("You entered: %v\n", charVar)
  fmt.Printf("You entered: '%v'\n\n", string(charVar))

    // Limiting by Characters
  fmt.Println("Please enter a single character (It can only be A, a, Z, or z) ")
  charVar = reader.ReadCharacterSet([]rune("AaZz"))
  fmt.Printf("You entered: %v\n", charVar)
  fmt.Printf("You entered: '%v'\n\n", string(charVar))

  // Getting strings, allowing the empty string ""
  fmt.Println("Please enter a string value ")
  stringVar = reader.ReadString(true)
  fmt.Printf("You entered: %v\n\n", stringVar)

  // Not allowing the empty string ""
  fmt.Println("Please enter a string value, it cannot be empty ")
  stringVar = reader.ReadString(false)
  fmt.Printf("You entered: %v\n\n", stringVar)

  // The Program Ends Here
}