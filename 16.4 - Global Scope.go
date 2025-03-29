/*
 * Author      : Luke Sathrum
 * Description : Examples of a global constant PI being used inside multiple
 *               functions.
 */

package main

import (
  "fmt"
  "math"
)
// Declare PI as a global constant
const kPi float64 = 3.14159

func area(radius float64) float64 {
  // Use our Global Constant
  return kPi * math.Pow(radius, 2);
}

func circumference(radius float64) float64 {
  return kPi * radius * 2
}

// Program Starts Here
func main() {
	// Use our global constant
  fmt.Println("For reference PI is defined as", kPi)

  // Output our Functions
  fmt.Println("The area of radius 5 is", area(5))
  fmt.Println("The circumference of radius 10 is", circumference(10))
}
