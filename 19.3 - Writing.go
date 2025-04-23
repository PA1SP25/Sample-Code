/*
 * Author      : Luke Sathrum
 * Description : Writing to files. Depends on story.txt
 */
package main

import (
	"fmt"
  "os"
)

func handleError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
	// Open a file to write to
  file, err := os.Create("write.txt")
  // Handle any errors
  handleError(err)
  // Close the file
  defer file.Close()

  // Write a string to the file
  n, err := file.WriteString("This is the first line\n")
  fmt.Println("Successfully outputed ", n, "bytes")

  // Can also write a byte slice
  bSlice := []byte{'H', 'E', 'L', 'L', 'O'}
  n, err = file.Write(bSlice)
  fmt.Println("Successfully outputed ", n, "bytes")

  // "Sync the file"
  file.Sync()

  // Open another file in append mode
  file2, err := os.OpenFile("story.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
  // Make sure it opened
  handleError(err)
  // Close the file
  defer file2.Close()

  // Lazy Writing, adding to our story
  _, _ = file2.WriteString("\nhair on each toe.")
}
