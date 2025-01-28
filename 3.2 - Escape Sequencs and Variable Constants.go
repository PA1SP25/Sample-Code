/*
 * Author      : Luke Sathrum
 * Description : Shows common escape sequences and how to create constants
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // \n - New Line
  fmt.Print("Let's enter a new line\nThis is a new line\n")

  // \t Horizontal Tab
  fmt.Print("Let's put a tab between this\tand this\n")

  // \\ Backslash
  // fmt.Println("This is NOT how we get a backslash (\) in a string")
  fmt.Println("This is how we get a backslash (\\) in a string")

  // \" Double Quote
  fmt.Println("Here is a double quote in a string: \"")
  // Incorrect Way
  // fmt.Println("Here is a double quote in a string: "")

  // \uhhhh Unicode - https://www.w3schools.com/charsets/ref_html_utf8.asp
  fmt.Println("\u2615")



  //  Declare a constant variable
  const kRate float32 = 6.9
  // What happens if we redeclare rate?
  // kRate = 7.0

  // Declare variables needed for constant example
  var deposit float32
  var new_balance float32

  // Get the deposit amount
  fmt.Print("Enter the amount of your deposit $")
  fmt.Scanf("%f", &deposit)

  // Calculate the balance based on rate and deposit  
  new_balance = deposit + (deposit * (kRate / 100))


  // Output the Results
  fmt.Println("The current rate is:", kRate)
  fmt.Printf("In one year, that deposit will grow to\n$%.2f, an amount worth waiting for.\n", new_balance)
  
  
  // This ends the program
}