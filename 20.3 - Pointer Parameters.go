/*
 * Author      : Luke Sathrum
 * Description : Pointers as Parameters
 */
package main

import (
	"fmt"
)

// notSneaky will NOT change the value of the parameter AFTER the call
func notSneaky(param int) {
  // Assigning the value 99 to the parameter
  param = 99
  fmt.Println("Inside function notSneaky param    ==", param)
}

// sneaky WILL change the value of the parameter AFTER the call
func sneaky(param *int) {
  // Assigning the value 99 to the parameter
  *param = 99
  fmt.Println("Inside function sneaky *param   ==", *param)
}

func main() {
	// Create a variable and give it a value
  variable := 100

  // Call notSneaky()
  fmt.Println("Before call to notSneaky, variable ==", variable)
  notSneaky(variable)
  fmt.Println("After call to notSneaky, variable  ==", variable, "\n")

  // Call sneaky()
  fmt.Println("Before call to sneaky, variable ==", variable)
  sneaky(&variable)
  fmt.Println("After call to sneaky, variable  ==", variable, "\n")
}
