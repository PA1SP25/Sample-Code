/*
 * Author      : Luke Sathrum
 * Description : Examples of slice functions append() and copy()
 */
package main

import "fmt"

func main() {
	// An empty slice
  var theSlice []string
  fmt.Println("theSlice", theSlice)
  fmt.Println("len(theSlice)", len(theSlice))
  fmt.Println("cap(theSlice)", cap(theSlice))
  fmt.Println()

  // Using Append
  theSlice = append(theSlice, "A", "B", "C")
  fmt.Println("theSlice", theSlice)
  fmt.Println("len(theSlice)", len(theSlice))
  fmt.Println("cap(theSlice)", cap(theSlice))
  fmt.Println()

  // You can create new slices using append
  theNewSlice := append(theSlice, "D", "E")
  fmt.Println("theSlice", theSlice)
  fmt.Println("len(theSlice)", len(theSlice))
  fmt.Println("cap(theSlice)", cap(theSlice))
  fmt.Println("theNewSlice", theNewSlice)
  fmt.Println("len(theNewSlice)", len(theNewSlice))
  fmt.Println("cap(theNewSlice)", cap(theNewSlice))
  fmt.Println()

  // copy, right and wrong

  // Wrong way to copy
  slice1 := []int{1, 2, 3}
  slice2 := slice1
  // Output the values in the slices
  fmt.Println("Wrong Way\nThe Slices")
  fmt.Println(slice1)
  fmt.Println(slice2)

  // Change a value in slice1
  slice1[0] = 10
  // Notice how it also changes slice2
  fmt.Println("The Slices")
  fmt.Println(slice1)
  fmt.Println(slice2)
  fmt.Println()

  // Correct Way
  slice3 := []int{1, 2, 3}
  slice4 := make([]int, len(slice3))
  copy(slice4, slice3)
  // Output the values in the slices
  fmt.Println("Correct Way\nThe Slices")
  fmt.Println(slice3)
  fmt.Println(slice4)

  // Change a value in slice1
  slice3[0] = 10
  // Notice how it also changes slice2
  fmt.Println("The Slices")
  fmt.Println(slice3)
  fmt.Println(slice4)
}
