/*
 * Author      : Luke Sathrum
 * Description : Creating sub-structures and initializing structures
 */

package main

import "fmt"

// Create a Date structure
type Date struct {
  month int
  day int
  year int
}

// Create a Structure to hold person information
type PersonInfo struct {
  height float64
  weight int
  birthday Date
}


// Program Starts Here
func main() {
  // Create and initialize a Date structure
  today := Date{1, 1, 1900}
  fmt.Printf("Today's date is %d/%d/%d\n\n", today.month, today.day, today.year)

  // Create a Person
  var person PersonInfo
  // Setup some of their data
  person.height = 6
  person.weight = 175
  // We are allowed to assign a structure's data to another structure
  person.birthday = today

  // Output all of the person's info
  fmt.Printf(`The person is %.2f feet tall and
weighs %d pounds and was born %d/%d/%d
`, person.height, person.weight, person.birthday.month, person.birthday.day, person.birthday.year)


}
