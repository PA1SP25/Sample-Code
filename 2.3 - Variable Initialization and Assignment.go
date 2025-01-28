/*
 * Author      : Luke Sathrum
 * Description : Assign values to variables, testing uninitialized
 */
package main

import "fmt"

func main() {
  // Declare some integers and initialize
  int1 := 0
  int2, int3 := 10, 0

  // Declare some floating-points, some initialized
  float1, float2 := 2.5, 3.3
  var float3 float32

  // Declare a string
  string1 := "This is a string"

  // Output
  fmt.Println(int1, int2, int3)
  fmt.Println(float1, float2, float3)
  fmt.Println("The output is:", string1)

  // Shorthand Assignment
  int1 += 5
  float1 += 5.2;

  fmt.Println("int:", int1)
  fmt.Println("float:", float1)


  // This ends the program
}