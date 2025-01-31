/*
 * Author      : Luke Sathrum
 * Description : Examples of output to the console
 */

package main

import "fmt"

// Program execution starts here
func main() {
  // A simple output to the user using Print
  fmt.Print("Hello reader.\nWelcome to Go!")

  // Println adds a newline at the end
  fmt.Println("Hello reader.\nWelcome to Go!")

  // Can separate by commas and add lots of types
  x := 12
  y := 7.7
  z := true
  fmt.Println("This", x, y, z, 10, 2.3, false, x - 10)
  

  // Printf using verbs
  // Integers
  fmt.Println("Integers")
  fmt.Printf("%v\n", x)
  fmt.Printf("%b\n", x)
  fmt.Printf("%o\n", x)
  fmt.Printf("%O\n", x)
  fmt.Printf("%d\n", x)
  fmt.Printf("%x\n", x)
  fmt.Printf("%X\n\n\n", x)

  // Floating-Points
  fmt.Println("Floating-Points")
  fmt.Printf("%v\n", y)
  fmt.Printf("%f\n", y)
  fmt.Printf("%g\n\n", y)

  // Floating-Points with Precision
  fmt.Println("Floating-Points with Precision")
  fmt.Printf("%f\n", y)
  fmt.Printf("%9f\n", y)
  fmt.Printf("%.2f\n", y)
  fmt.Printf("|%9.2f|\n", y)
  fmt.Printf("|%-9.2f|\n", y)
  fmt.Printf("%9.f\n\n", y)

  // More Practice with Precision
  fmt.Printf("You owe me $%.2f. Please pay promptly!\n\n", y)
  
  // Strings
  fmt.Println("Strings")
  str := "World"
  fmt.Printf("Hello %v!\n", str)
  fmt.Printf("Hello %s!\n\n", str)


  // Booleans
  fmt.Println("Booleans")
  b := true
  fmt.Printf("%v\n", b)
  fmt.Printf("%t\n\n", false)

  // All Together
  fmt.Printf("There were %d foxes who bought a $%.2f box. Is this story %t?\n", x, y, b)
  fmt.Printf(`There were %d foxes who bought a $%.2f box.
Is this story %t?
`, x, y, b)


  // The Program Ends Here
}