/*
 * Author      : Luke Sathrum
 * Description : Examples of creating pointers and using * and & operators.
 */
package main

import (
	"fmt"
)

func main() {
	// Creating some pointers
  var p1, p2 *int
  // Ouput their "zero" values
  fmt.Println(p1, p2, "\n")
  // Creating some integer variables
  var v1, v2 int = 1, 2

  // Point p1 to variable v1
  p1 = &v1
  // Point p2 to variable v2;
  p2 = &v2

  // Output the values of v1 and *p1
  fmt.Println("v1: ", v1)
  fmt.Println("*p1:", *p1, "\n")
  
  // Output the values of v2 and *p2
  fmt.Println("v2: ", v2)
  fmt.Println("*p2:", *p2, "\n")

  // We can also output the memory addresses of p1 and p2
  // As well as v1 and v2
  fmt.Println("p1: ", p1)
  fmt.Println("&v1:", &v1, "\n")

  fmt.Println("p2: ", p2)
  fmt.Println("&v2:", &v2, "\n")

  // If you want to get real tricky you can output the memory address of
  // the pointers
  fmt.Println("&p1:", &p1)
  fmt.Println("&p2:", &p2)
}
