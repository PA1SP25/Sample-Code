/*
 * Author      : Luke Sathrum
 * Description : Examples of Operator Precedence
 */

package main

import "fmt"

func main() {
	// Declare our Variables
	var (
		a = 20
		b = 10
		c = 15
		d = 4
		e int
	)

	// Arithmetic Precedence
	// Run the first example (20 + 10) * 15 / 5
	e = (a + b) * c / d
	fmt.Println("Value of (a + b) * c / d is    :", e)

  // Run the second example ((20 + 10) * 15) / 5
  // This is the preferred style
  e = ((a + b) * c) / d;
  fmt.Println("Value of ((a + b) * c) / d is  :", e)

  // Run the third example (20 + 10) * (15 / 5)
  e = (a + b) * (c / d);
  fmt.Println("Value of (a + b) * (c / d) is  :", e)

  // Run the fourth example 20 + (10 * 15) / 5
  e = a + (b * c) / d;
  fmt.Println("Value of a + (b * c) / d is    :", e)

// Run the fifth example 20 + 10 * 15 / 5
  e = a + b * c / d;
  fmt.Println("Value of a + b * c / d is      :", e)

  // Combined Precedence
  x := 5
  y := 2
  if x + 1 < -2 || x + 1 > 3 {
    fmt.Println("Combined Precedence was TRUE")
  } else {
    fmt.Println("Combined Precedence was FALSE")
  }

  z := y + 1 < -2 || x + 2 * 3 >= 12
  fmt.Println("z is :", z)

}
