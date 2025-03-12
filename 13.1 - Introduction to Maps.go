/*
 * Author      : Luke Sathrum
 * Description : Examples of Maps in Go
 */

package main

import "fmt"

func main() {
	// Create a map of runes -> strings
	values := make(map[rune]string)

	// Add some key/value pairs
	values['A'] = "Alpha"
	values['B'] = "Bravo"
	values['C'] = "Charlie"
	values['D'] = "Delta"

	// Output some information about the Map
	fmt.Println(values)
	fmt.Println(len(values))

	// Change a Map value
	values['A'] = "Animal"
	fmt.Println(values)

	// Accessing a non-existant key
	fmt.Printf("%s was located at key 'Z'\n", values['Z'])

	// Handling BOTH values that accessing provides
	val, res := values['Z']
	fmt.Printf("val: %s, res: %t\n", val, res)

	// We can also assign in an if
	if val, res := values['Z']; res {
		fmt.Println(val)
	} else {
      fmt.Println("That key does not exist in the Map")
  }
}