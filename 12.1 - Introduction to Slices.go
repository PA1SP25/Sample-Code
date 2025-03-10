/*
 * Author      : Luke Sathrum
 * Description : Examples of slices in Go
 */
package main

import "fmt"

func main() {
	// Create an "empty" slice
  var temps[]float64

  // Output some info about temps
  fmt.Println("----- var temps[]float64 -----")
  fmt.Println("temps:", temps)
  fmt.Println("len(temps): ", len(temps))
  // You can't access anything since there is no size
  // temps[0] = 5
  // fmt.Println(temps[0])

  // Create a slice of a given size
  values := make([]int, 5)
  // Output some info about temps
  fmt.Println("\n----- values := make([]int, 5) -----")
  fmt.Println("values:", values)
  fmt.Println("len(values): ", len(values))
  // We can set using indicies 0 - 4
  values[3] = 5
  fmt.Println("values:", values)
  // We can loop through as well, outputting as comma separated list
  for i := 0; i < len(values) - 1; i++ {
    fmt.Printf("%d, ", values[i])
  }
  fmt.Println(values[len(values) - 1])

  // Create a slice of a given size and capacity
  flags := make([]bool, 3, 6)
  // Output some info about flags
  fmt.Println("\n----- flags := make([]bool, 3, 6) -----")
  fmt.Println("flags:", flags)
  fmt.Println("len(flags): ", len(flags))
  fmt.Println("cap(flags): ", cap(flags))
  // We can set using indicies 0 - 2
  flags[2] = true
  fmt.Println("flags:", flags)
  // Currently cannot use indicies 3 - 5
  // flags[5] = true
  // fmt.Println("flags:", flags)
}
