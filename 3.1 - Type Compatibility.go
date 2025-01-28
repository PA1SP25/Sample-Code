/*
 * Author      : Luke Sathrum
 * Description : Shows compatibility between types
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Declare some variables
  var intVar int = 1
  var floatVar float32 = 1.5
  var stringVar string = "Hello"
  
  // This doesn't work. 2.99 is a double and we are storing it as an int
  // intVar = 2.99
  // fmt.Println("intVar:", intVar)
  
  // This is OK
  floatVar = 2
  fmt.Println("floatVar:", floatVar)

  // This doesn't work
  stringVar = 5
  stringVar = floatVar
  
  // This ends the program
}