/*
 * Author      : Luke Sathrum
 * Description : Examples of local variables
 */

package main

import "fmt"

// Function Declaration for Dummy
func dummy(one int) {
  // Note that one ^^^ is a local variable of the function Dummy
  var two int = 5
  fmt.Println("dummy() one + dummy() two =", one + two)
}

// Program Starts Here
func main() {
	// Setup some local variables
  var (
    one = 1.0
    two = 2
  )
  
  // Output one plus two
  fmt.Println("local one + local two =", one + float64(two))
  // Get output of dummy()
  dummy(two);
}
