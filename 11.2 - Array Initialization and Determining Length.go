/*
 * Author      : Luke Sathrum
 * Description : Initialize an array using {} and
 *               find the size using len()
 */
package main

import "fmt"

func main() {
  // Create and initialize an array
  children := [2]string{"Odessa", "Santos"}
  fmt.Println("The oldest   child's name is", children[0])
  fmt.Println("The youngest child's name is", children[1])

  // Create another array for ages
  var childrenAges[2]int
  // Loop through the age array and get the ages
  for i := 0; i < len(childrenAges); i++ {
    fmt.Printf("Enter the age of %s: ", children[i])
    fmt.Scan(&childrenAges[i])
  }

  // Output the ages
  fmt.Print("The ages were: ")
  for i := 0; i < len(childrenAges); i++ {
    fmt.Print(childrenAges[i])
    // Don't output a comma at the end
    if i < len(childrenAges) - 1 {
      fmt.Print(", ")
    }
  }
  fmt.Println()
}