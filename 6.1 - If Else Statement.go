/*
 * Author      : Luke Sathrum
 * Description : Using if/else statements calculate a user's pay, including
 *               any overtime.
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // Declare Variables
  var hours, grossPay, grossOvertime float64
  rate := 30.0

  // Prompt for work hours
  fmt.Print("How many hours did you work this week? ")
  fmt.Scan(&hours)
  
  // if-else statement to compute overtime if worked
  if hours > 40 {
    grossPay = (rate * 40) + (1.5 * rate * (hours - 40))
    fmt.Println("You worked a long week!")
  } else {
    grossPay = rate * hours
    fmt.Println("Have you ever considered working overtime?")
  }
  // Output the results of the if-else
  fmt.Printf("Your gross pay is: $%.2f\n", grossPay)

  // Omitting the else
  if hours > 40 {
    grossOvertime = (1.5 * rate * (hours - 40))
    fmt.Println("You worked a long week")
    fmt.Printf("Your overtime pay this week was: $%.2f\n", grossOvertime)
  }


  // The Program Ends Here
}