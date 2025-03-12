/*
 * Author      : Luke Sathrum
 * Description : Sorting Slices
 */
package main

import (
  "fmt"
  "sort"
)

// Checks if a Slice is sorted and outputs to the user the results
func checkIfSorted(slice []int) {
  if sort.IntsAreSorted(slice) {
    fmt.Println("The slice is sorted")
  } else {
    fmt.Println("The slice is NOT sorted")
  }
}

// Outputs a Slice to the user in a nice format
func outputPrettyIntSlice(slice []int) {
  for i := 0; i < len(slice) - 1; i++ {
    fmt.Printf("%d, ", slice[i])
  }
  fmt.Println(slice[len(slice) - 1])
}

func main() {
  // Create a slice and initialize
  values := []int{4, 3, 2, 1}

  // Output the slice
  outputPrettyIntSlice(values)

  // Is the slice sorted?
  checkIfSorted(values)

  // Sort the slice
  sort.Ints(values)

  // Output the slice
  outputPrettyIntSlice(values)

  // Is the slice sorted?
  checkIfSorted(values)
}