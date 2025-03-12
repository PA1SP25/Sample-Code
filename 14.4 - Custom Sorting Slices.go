/*
 * Author      : Luke Sathrum
 * Description : Custom Sorting Slices
 */
package main

import (
  "fmt"
  "sort"
)

// Outputs an Integer Slice to the user in a nice format
func outputPrettyIntSlice(slice []int) {
  for i := 0; i < len(slice) - 1; i++ {
    fmt.Printf("%d, ", slice[i])
  }
  fmt.Println(slice[len(slice) - 1])
}

// Outputs a String Slice to the user in a nice format
func outputPrettyStringSlice(slice []string) {
  for i := 0; i < len(slice) - 1; i++ {
    fmt.Printf("%s, ", slice[i])
  }
  fmt.Println(slice[len(slice) - 1])
}

func main() {
  // Create a slice and initialize
  values := []int{1, 4, 2, 3}

  // Output the slice
  outputPrettyIntSlice(values)

  // Custom Sort the Ints from 
  sort.Slice(values, func(prev, next int) bool {
    return values[prev] > values[next]
  })

  // Output the slice
  outputPrettyIntSlice(values)

  // Create a slice of strings and initialize
  words := []string{"foobar", "foos", "bar"}

  // Output the slice
  outputPrettyStringSlice(words)

  // Custom sort the strings by length
  sort.Slice(words, func(prev, next int) bool {
    return len(words[prev]) < len(words[next])
  })

  // Output the slice
  outputPrettyStringSlice(words)
}

