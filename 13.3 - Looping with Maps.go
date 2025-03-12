/*
 * Author      : Luke Sathrum
 * Description : Examples of looping through maps
 */
package main

import "fmt"

func main() {
  // Create and initialize a map
  animals := map[int]string{
  	90: "Dog",
  	91: "Cat",
  	92: "Cow",
  	93: "Bird",
  	94: "Rabbit",
  }
  // "Ugly" Print the Map
  fmt.Println(animals)

  // "Pretty" Print the Map
  for key, value := range animals {
    fmt.Printf("%-7s is located at key %d\n", value, key)
  }

  fmt.Println()
  fmt.Println()
  // Incorrect Copy
  animals2 := make(map[int]string)
  animals2 = animals
  // At this point, both maps are really the same map
  fmt.Println("Animals: ", animals)
  fmt.Println("Animals2:", animals2)
  fmt.Println()
  
  // Change animals map (which also changes animals2)
  animals[91] = "Lion"
  fmt.Println("Animals: ", animals)
  fmt.Println("Animals2:", animals2)
  fmt.Println()
  
  // Actually Make a Copy
  animals2 = make(map[int]string)
  // Loop and copy the value
  for key, value := range animals {
    animals2[key] = "2 " + value + "s"
  }

  fmt.Println("Animals: ", animals)
  fmt.Println("Animals2:", animals2)
}