/*
 * Author      : Luke Sathrum
 * Description : Converting Pseudocode to Go Code
 */
 
// Package Declaration
package main
// Include the built-in "fmt" package
import "fmt"

// Program Execution Starts Here
func main() {
  // Setup Storage
  var num1 int
  var num2 int
  var result int
  // Get number 1 from user and store it
  fmt.Println("Please enter the first number: ")
  fmt.Scanf("%d", &num1)
  // Get number 2 from user and store it
  fmt.Println("Please enter the second number: ")
  fmt.Scanf("%d", &num2)
  // Add number 1 and number 2 and store the result
  result = num1 + num2
  // Output the result to the user
  fmt.Printf("The Sum of %d and %d is: %d\n", num1, num2, result )
}




  // Get number 1 from user and store it
  // Get number 2 from user and store it
  // Add number 1 and number 2 and store the result
  // Output the result to the user