/*
 * Author      : Luke Sathrum
 * Description : Examples of the increment and decrement operators
 */

package main

import "fmt"

// Program execution starts here
func main() {
  x, y := 0, 10

  // These will not work
  // --x
  // (x + 2)++
  // --(y + 2)
  // 5++

  fmt.Println("x:", x, "y:", y)
  x++
  y--
  fmt.Println("x:", x, "y:", y)

  // The Program Ends Here
}