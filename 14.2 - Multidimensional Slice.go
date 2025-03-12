/*
 * Author      : Luke Sathrum
 * Description : Using a multidimensional slice to hold students' quiz scores.
 */
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Create a variable to hold # of quizes
  var numberOfQuizzes int

  // Create the start of a 2D slice (the rows)
  scores := make([][]int, 3)

  // Loop through and fill the rows with data
  for iRow := 0; iRow < len(scores); iRow++ {
    fmt.Printf("How many scores for Student %d? ", iRow + 1)
    fmt.Scan(&numberOfQuizzes)
    // Create an "Inner" slice based on user input
    scores[iRow] = make([]int, numberOfQuizzes)
    // Loop through the "Inner" slice and fill it up
    fmt.Print("Please enter the scores, space separated: ")
    for iCol := 0; iCol < numberOfQuizzes; iCol++ {
      fmt.Scan(&scores[iRow][iCol])
    }
  }
  fmt.Println()

  // Output the 2D Slice (ugly)
  fmt.Println(scores)
  fmt.Println()
  
  
  // Loop through and output the 2D slice
  // To output a multidimensional slice, we need a for loop for each dimension
  // len(scores) is the # of rows
  // len(scores[iRow] is the # of columns for that row)
  // This outside loop handles each student one-by-one
  for iRow := 0; iRow < len(scores); iRow++ {
    fmt.Printf("Student %d's quiz scores\n%s\n", iRow + 1, strings.Repeat("-", 30))
    // This inside loop handles each individual's quiz one-by-one
    for iCol := 0; iCol < len(scores[iRow]); iCol++ {
      fmt.Printf("Quiz %d: %d\n", iCol + 1, scores[iRow][iCol])
    }
    fmt.Printf("%s\n", strings.Repeat("-", 30))
  }
  fmt.Println()

  // You can also create and initialize a slice at the same time
  synonyms := [][]string{
    {"hello", "hi", "howdy"},
    {"foo"},
    {"goodbye", "so long"},
  }

  // Output the initialized 2D slice in a table like format using for range
  for _, synonymSlice := range synonyms {
    for _, synonym := range synonymSlice {
      fmt.Printf("%-10s", synonym)
    }
    fmt.Println()
  }
}

