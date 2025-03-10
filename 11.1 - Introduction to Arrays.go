/*
 * Author      : Luke Sathrum
 * Description : Our first use of arrays. We will get 5 scores and see how much
 *               each differs from the highest score.
 */
package main

import "fmt"


// We are going to read in five scores and show how much
// each differs from the highest score
func main() {
  // We need an array
  var scores[5]int
  // We need to store the highest score
  var highest int

  // Prompt for the first score
  fmt.Print("Enter 5 Scores, space separated: ")
  // Get the first score
  fmt.Scan(&scores[0])
  // Since it's the only value we have seen so far, it's the "highest"
  highest = scores[0]

  // Write a loop to get the rest of the scores
  for i := 1; i < 5; i++ {
    // Get the next score
    fmt.Scan(&scores[i])
    // Check to see if the new score is "higher", if it is... store
    if scores[i] > highest {
      highest = scores[i]
    }
  }

  // At this point our array is full and we can see the values
  fmt.Println(scores)

  // We can also output the highest score
  fmt.Println("The highest score is", highest)

  // Output the diffence of each score to the highest using a loop
  for i := 0; i < 5; i++ {
    fmt.Printf("%-3d off by %d\n", scores[i], highest - scores[i])
  }
}