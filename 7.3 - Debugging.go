/*
 * Author      : Luke Sathrum
 * Description : Debugging Syntax, Run-Time, and Logical Errors.
 *               NOTE: This code will not compile as is.
 */

package main

import "fmt"

func main() {
	// Syntax Errors
  // After compile you will get compiler errors
  var x

  flot Floatvar

  something
  
  if (true) {

  var unused int

  // Run-Time Error
  y := 10 / 0

  // Logical Error
  fmt.Println("8 * 5 =", 8 + 5)

}
