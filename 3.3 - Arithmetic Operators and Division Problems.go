/*
 * Author      : Luke Sathrum
 * Description : Arithmetic Operators with different types
 *               and integer and double division
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // First I want to show you arithmetic operators with different types
 
  // int + int = int
  fmt.Println("int 6 + int 9 -> int", 6 + 9)
 
  // float + float = float
  fmt.Println("float 6.1 + float 9.2 -> float", 6.1 + 9.2)
 
  // float + int = float
  fmt.Println("float 6.1 + int 9 -> float", 6.1 + 9)

  // int + float = float
  fmt.Println("int 6 + float 9.2 -> float", 6 + 9.2, "\n\n")

  // Float Division
  fmt.Println("float 6.0 / float 9.0 =", 6.0 / 9.0)
  // Integer Division
  fmt.Println("int 6 / int 9 =", 6 / 9)
  // Remainer Division (Mod)
  fmt.Println("int 6 % int 9 =", 6 % 9)
  // To get a float result we need to make one of the numbers a float
  fmt.Println("int 6 / float 9.0 =", 6 / 9.0)

  // The Program Ends Here
}