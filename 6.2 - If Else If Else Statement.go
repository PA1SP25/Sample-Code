/*
 * Author      : Luke Sathrum
 * Description : Modifying our user's pay scenario to include if else-if else
 *               statements.
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Declare Variables
  var hours, grossPay float64
  rate := 30.0

  // Prompt for work hours
  fmt.Print("How many hours did you work this week? ")
  fmt.Scan(&hours)
  
  // if-else statement to compute overtime if worked
  if hours > 40 {
    grossPay = (rate * 40) + (1.5 * rate * (hours - 40))
  } else if hours > 0 {
    grossPay = rate * hours
  } else {
    fmt.Println("You did not work this week!");
  }
  // Output the results of the if-else
  fmt.Printf("Your gross pay is: $%.2f\n", grossPay)

  // The Program Ends Here
}