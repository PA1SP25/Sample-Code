/*
 * Author      : Luke Sathrum
 * Description : Examples of map initialization and key deletion
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
  
  // Output the Map
  fmt.Println(animals)
  
  // Remove key 92 from the map
  delete(animals, 92)
  
  // Output the Map Again
  fmt.Println(animals)
}