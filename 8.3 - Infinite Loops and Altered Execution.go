/*
 * Author      : Luke Sathrum
 * Description : Infinite Loops and Altering Flow
 */

package main

import "fmt"

func main() {
	// Infinite Loops
  // Simple Example
  // for {
  //   fmt.Println("Infinite")
  // }

  // Output even numbers from 2 - 10
  for i := 2; i != 12; i += 2 {
    fmt.Println(i)
  }

  // Now the odd numbers from 1 - 11
  // for i := 1; i != 12; i += 2 {
  //   fmt.Println(i)
  // }

  // Altering Flow (break)
  // x := 0
  // for {
  //   fmt.Println("Hello")
  //   if (x == 10) {
  //     break
  //   }
  //   x++
  // }

  // Altering Flow (continue)
  // for i := 1; i <= 10; i++ {
  //   if (i % 2 == 0) {
  //     continue
  //   }
  //   fmt.Println(i)
  // }

}
