/*
 * Author      : Luke Sathrum
 * Description : Examples of slice initialization
 */
package main

import "fmt"

func main() {
  // Create and Initialize a slice
	theSlice := []bool{true, false, true}
  fmt.Println(theSlice)
  fmt.Println()


  // Create an Array
  theArray := [3]string{"A", "B", "C"}

  // Create Slices from the array
  slice1 := theArray[1:2]
  slice2 := theArray[1:]
  slice3 := theArray[:2]
  slice4 := theArray[:]

  // Output the Slices
  fmt.Println("Slice    Length   Capacity")
  fmt.Println(slice1, "       ", len(slice1), "       ", cap(slice1))
  fmt.Println(slice2, "     ", len(slice2), "       ", cap(slice2))
  fmt.Println(slice3, "     ", len(slice3), "       ", cap(slice3))
  fmt.Println(slice4, "   ", len(slice4), "       ", cap(slice4))
  fmt.Println()
  
  // Currently the slices are pieces of theArray
  theArray[1] = "Z"
  fmt.Println(theArray)
  fmt.Println(slice1)
  fmt.Println(slice2)
  fmt.Println(slice3)
  fmt.Println(slice4)
  fmt.Println()

  // Even though the slice is from theArray, it's indicies are it's own
  fmt.Println(theArray)
  slice1[0] = "S"
  fmt.Println(theArray)
  fmt.Println()

  // Our for range loops also work with slices
  for _, val := range slice3 {
    fmt.Println(val)
  }
}
