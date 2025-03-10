/*
 * Author      : Luke Sathrum
 * Description : Examples of Call-by-Value and Call-by-Reference
 */

package main

import "fmt"

/*
 * Changes the value of the parameter and outputs it
 */
func callByValue(callByValue string) {
  callByValue = "One"
  fmt.Println("Inside CallByValue():", callByValue)
}

/*
 * Changes the value of the parameter and outputs it
 */
func callByReference(callByReference *string) {
  *callByReference = "Two"
  fmt.Println("Inside CallByReference()", *callByReference)
}

/*
 * Can Mix Parameters
 */
func callBy(callByValue string, callByReference *string) {
  callByValue = "One"
  *callByReference = "Three"
  fmt.Println("Inside CallBy():", callByValue)
  fmt.Println("Inside CallBy()", *callByReference)
}

/*
 * Get's an integer from the user (by-reference)
 */
func input(val *int) {
  fmt.Print("Enter a whole number: ")
  // Usually we add an & with Scan, but not here
  fmt.Scan(val)
}

func main() {
  word, word2 := "Hello", "Goodbye"
  fmt.Println("The value before a call to callByValue():", word)
  callByValue(word)
  fmt.Println("The value after a call to callByValue():", word)
  fmt.Println()

  fmt.Println("The value before a call to callByReference():", word)
  // Need to add an & when by-reference
  callByReference(&word)
  fmt.Println("The value after a call to callByReference():", word)
  fmt.Println()

  fmt.Println("The value before a call to callBy:", word, word2)
  // Need to add an & when by-reference
  callBy(word, &word2)
  fmt.Println("The value after a call to callBy():", word, word2)
  fmt.Println()

  // Use Call-By-Reference to Get input
  var x int
  input(&x)
  fmt.Println("You entered:", x);
}
