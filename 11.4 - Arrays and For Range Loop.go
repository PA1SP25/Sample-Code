/*
 * Author      : Luke Sathrum
 * Description : Examples of for range loops with arrays
 */
package main

import "fmt"

func main() {
  // Create an array to hold cities as strings
  var cities[5]string

  // Input city information into our array
  // Have to use cities[i] to store
  for i, _ := range cities {
    fmt.Printf("Enter City #%d: ", i + 1)
    fmt.Scan(&cities[i])
  }

  // Output the cities using a for range loop
  fmt.Println("\nThe cities entered were\n-----------------------")
  for _, city := range cities {
    fmt.Println(city)
  }
}