/*
 * Author      : Luke Sathrum
 * Description : Short Circuit Evaluation
 */

package main

import "fmt"

func main() {
	// Declare our Variables
	pieces := 10
  students := 0
  fmt.Print("How many students? ")
  fmt.Scan(&students)

  if (pieces / students) >= 2 {
    fmt.Println("Each student may have two pieces.")
  } else {
    fmt.Println("There is not enough candy for everyone! :(")
  }

  // A Better way to write it
  // if (students > 0) && (pieces / students) >= 2 {
  //   fmt.Println("Each student may have two pieces.")
  // } else {
  //   fmt.Println("There is not enough candy for everyone! :(")
  // }

}
