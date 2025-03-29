/*
 * Author      : Luke Sathrum
 * Description : Defining and Using Structures
 */

package main

import "fmt"

// Create a Car Structure w/ 5 fields of various types
type Car struct {
  year int
  doors int
  horsePower float64
  cmake string
  model string
}

// Program Starts Here
func main() {
  // Create a Car
  var myCar, yourCar Car

  // Setup myCar's Data
  myCar.year = 2005
  myCar.doors = 4
  myCar.horsePower = 255
  myCar.cmake = "Honda"
  myCar.model = "Pilot"

  // Output the data in 2 ways
  fmt.Println(myCar)
  fmt.Printf("%#v\n", myCar)

  // Get yourCar Data from the User
  fmt.Print("What is the year of your car? ")
  fmt.Scan(&yourCar.year)
  fmt.Print("How many doors does your car have? ")
  fmt.Scan(&yourCar.doors)
  fmt.Print("How much horsepower does your car have? ")
  fmt.Scan(&yourCar.horsePower)
  fmt.Print("What is the make of your car? ")
  fmt.Scan(&yourCar.cmake)
  fmt.Print("What is the model of your car? ")
  fmt.Scan(&yourCar.model)

  // Pretty Formatted Output
  fmt.Println("\n")
  fmt.Printf("            %-15s%-15s\n", "My Car", "Your Car")
  fmt.Println("-----------------------------------------")
  fmt.Printf("Year:       %-15d%-15d\n", myCar.year, yourCar.year)
  fmt.Printf("Doors:      %-15d%-15d\n", myCar.doors, yourCar.doors)
  fmt.Printf("Horsepower: %-15.2f%-15.2f\n", myCar.horsePower, yourCar.horsePower)
  fmt.Printf("Make:       %-15s%-15s\n", myCar.cmake, yourCar.cmake)
  fmt.Printf("Model:      %-15s%-15s\n", myCar.model, yourCar.model)
}