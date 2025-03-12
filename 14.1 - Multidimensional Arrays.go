/*
 * Author      : Luke Sathrum
 * Description : Using a multidimensional array to hold students' quiz scores.
 */
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Create and initialize a multidimenstional array of 4 rows and 3 columns
  scores := [4][3]int{
  // Quiz 1, Quiz 2, Quiz 3  
    {10,     5,      7},  // Student 1
    {1,      2,      3},  // Student 2
    {9,      9,      8},  // Student 3
    {8,      7,      5},  // Student 4
  }

  // Change student 2's quiz 3 score
  scores[1][2] = 10

  // Output the entire 2D array
  fmt.Println(scores)

  // Output student 2's quiz scores
  fmt.Println(scores[1])
  
  // Loop through and output the 2D array
  // To output a multidimensional array, we need a for loop for each dimension
  // len(scores) is the # of rows
  // len(scores[iRow] is the # of columns)
  // This outside loop handles each student one-by-one
  for iRow := 0; iRow < len(scores); iRow++ {
    fmt.Printf("Student %d's quiz scores\n%s\n", iRow + 1, strings.Repeat("-", 30))
    // This inside loop handles each individual's quiz one-by-one
    for iCol := 0; iCol < len(scores[iRow]); iCol++ {
      fmt.Printf("Quiz %d: %d\n", iCol + 1, scores[iRow][iCol])
    }
    fmt.Printf("%s\n", strings.Repeat("-", 30))
  }
}

