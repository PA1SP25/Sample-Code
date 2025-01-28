/*
 * Author      : Luke Sathrum
 * Description : Shows how to correctly and incorrectly write identifiers
 */
package main

import "fmt"

func main() {
  // Incorrect
  // var 6num int
  // Correct
  // var num6 int

  // Incorrect
  // var -num int
  // Correct
  // var num int

  // Incorrect
  // var myfirst.c int
  // Correct
  // var myFirst int

  // Case Sensitive;
  // var rate int = 0
  // var RATE int = 1
  // var Rate int = 2
  // fmt.Println("rate: ", rate)
  // fmt.Println("RATE: ", RATE)
  // fmt.Println("Rate: ", Rate)


  // Reserved Words
  // var var int
  // var float32 int
  // var string int

  // Bad Naming
  // var x, y, z int = 0, 2, 3
  // x = y * z
  // fmt.Println("x: ", x)
  // Good Naming
  // var distance, speed, time int = 0, 2, 3
  // distance = speed * time
  // fmt.Println("distance: ", distance)

  // This ends the program

}