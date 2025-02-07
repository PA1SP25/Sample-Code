/*
 * Author      : Luke Sathrum
 * Description : Using constants and switch statements to calculate
 *               a toll.
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Create a group of constants to help with our switch (and documentation)
  const (
    kCar = 1
    kBus = 2
    kTruck = 3
  )
  // Declare Variables
  var vehicleClass int
  var toll float64

  // Get input
  fmt.Printf("Enter vehicle class (%d = Car, %d = Bus, %d = Truck): ", kCar, kBus, kTruck)
  fmt.Scan(&vehicleClass)

  // Make decisions based on vehicleClass
  switch (vehicleClass) {
    case kCar:
      fmt.Println("Passenger Car.")
      toll = 0.50
    case kBus:
      fmt.Println("Bus.")
      toll = 1.50
    case kTruck:
      fmt.Println("Truck.")
      toll = 2.00
    default:
      fmt.Println("Unknown vehicle class!")
      toll = 5.00
  }

  // Output to the user the toll
  fmt.Printf("The toll is $%.2f\n", toll);

/*
    You can write a switch as an if/else (and vice versa)
    if vehicle_class == kCar {
      fmt.Println("Passenger Car.")
      toll = 0.50
    } else if vehicle_class == kBus {
      fmt.Println("Bus.")
      toll = 1.50
    } else if vehicle_class == kTruck {
      fmt.Println("Truck.")
      toll = 2.00;
    } else {
      fmt.Println("Unknown vehicle class!")
      toll = 5.00
    }

    fmt.Printf("The toll is $%.2f\n", toll);
*/
  // The Program Ends Here
}