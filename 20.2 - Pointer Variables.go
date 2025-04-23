/*
 * Author      : Luke Sathrum
 * Description : Examples of creating "nameless" variables
 */
package main

import (
	"fmt"
)

func main() {
	// Creating some pointers
  var p1, p2 *int

  // Assign p1 to a "nameless" integer
  p1 = new(int)

  // Give the "nameless" variable the value 42
  *p1 = 42

  // Point p2 to the smae place as p1
  p2 = p1

  // Check the values
  fmt.Println("*p1 ==", *p1)
  fmt.Println("*p2 ==", *p2, "\n")

  
  // Assign 53 to whatever p2 is pointing to
  *p2 = 54

  // Check the values
  fmt.Println("*p1 ==", *p1)
  fmt.Println("*p2 ==", *p2, "\n")

  // Create another "nameless" variable
  p1 = new(int)

  // Check the values
  fmt.Println("*p1 ==", *p1)
  fmt.Println("*p2 ==", *p2, "\n")

  // Assign 88 to whatever p1 is pointing to
  *p1 = 88

  // Check the values
  fmt.Println("*p1 ==", *p1)
  fmt.Println("*p2 ==", *p2, "\n")
}
