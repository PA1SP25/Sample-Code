/*
 * Author      : Luke Sathrum
 * Description : Examples of type conversion
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Setup some variables
  x, y, resultInt, resultFloat := 9, 2, 0, 0.0

  // Do integer division
  resultInt = x / y

  // Do float division
  // resultFloat = x / y
  // resultFloat = float64(x) / y
  resultFloat = float64(x) / float64(y)

  // Output
  fmt.Println("The result of integer division was:", resultInt)
  fmt.Println("The result of float division was:", resultFloat)

  // Let's convert a float to an integer
  var num int
  num2 := 4.9
  num = int(num2)
  fmt.Println("The value of num is:", num)

  // The Program Ends Here
}