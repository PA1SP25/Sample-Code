/*
 * Author      : Luke Sathrum
 * Description : Initialize an array using {} and
 *               find the size using len()
 */
package main

import "fmt"

func main() {
  // Create an array
  a := [6]int{10, 20, 30, 40, 50, 60}

  // Output information about the array
  for i := 0; i < len(a); i++ {
    fmt.Printf("Value %d, at location %d, is located at memory location %p\n", a[i], i, &a[i])
  }
  fmt.Println()

  // Syntax Error (Out-of-Range)
  //fmt.Println(a[6])

  // Runtime Error (Out-of-Range)
  for i := 0; i <= len(a); i++ {
    fmt.Println(a[i])
  }
}