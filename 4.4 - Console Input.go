/*
 * Author      : Luke Sathrum
 * Description : Shows hows to get user input
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Setup some Variables
  var num1, num2 int

  // Prompt the user and get some input
  fmt.Print("Please enter a whole number: ")
  fmt.Scan(&num1)
  fmt.Println("You entered", num1)

  // Get multiple numbers from the user
  fmt.Print("Please enter 2 integer numbers separated by a space: ")
  fmt.Scan(&num1, &num2)
  fmt.Println("You entered", num1, num2)

  // Scanf Example
  fmt.Print("Please enter another whole number: ")
  fmt.Scanf("%d", &num1)
  fmt.Println("You entered", num1)


  // The Program Ends Here
}