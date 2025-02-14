/*
 * Author      : Luke Sathrum
 * Description : Advanced for loops
 */

package main

import "fmt"

func main() {
	// A for loops that counts
  for number := 10; number >= 0; number-- {
    fmt.Printf("T minus %d seconds and counting...\n", number)
  }

  // We can use if statements in loops
  // for number := 10; number >= 0; number-- {
  //   if number == 0 {
  //     fmt.Println("BLASTOFF!!!")
  //   } else if number == 1 {
  //     fmt.Printf("T minus %2d second  and counting...\n", number)
  //   } else {
  //     fmt.Printf("T minus %2d seconds and counting...\n", number)

  //   }
  // }

  // What happens if we forget to create our variable before we use it?
  // for number = 10; number >= 0; number-- {
  //   fmt.Printf("T minus %d seconds and counting...\n", number)
  // }
}