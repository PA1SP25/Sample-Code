/*
 * Author      : Luke Sathrum
 * Description : Opening and Reading files into a buffer. Depends on rhyme.txt
 */
package main

import (
	"fmt"
  "os"
  "bufio"
)

func handleError(err error) {
  if err != nil {
    panic(err)
  }
}

func main() {
	// Open a file
  file, err := os.Open("rhyme.txt")
  // Make sure the file opened
  handleError(err)
  // Close the file
  defer file.Close()

  // Create a reader from the opened file
  reader := bufio.NewReader(file)

  // How many bytes can be read in the buffer?
  byteSlice, err := reader.Peek(5)
  // Handle an error
  handleError(err)
  // Output the byteSlice as a string
  fmt.Println(string(byteSlice))

  // How many bytes can be read in the buffer?
  // byteSlice, err = reader.Peek(500)
  // Handle an error
  // handleError(err)
  // Output the byteSlice as a string
  // fmt.Println(string(byteSlice))

  // Read the next byte in the buffer
  char, err := reader.ReadByte()
  // Handle an error
  handleError(err)
  // Output the byte as a string
  fmt.Println(string(char))

  // Read the rest of the "line", as a byte slice
  byteSlice, err = reader.ReadBytes('\n')
  handleError(err)
  fmt.Println(string(byteSlice))

  // Read the next line as a string
  str, err := reader.ReadString('\n')
  handleError(err)
  fmt.Println(str)

  // Scanners are nice for easy reading of a line (it defaults to read a line)
  // We need to close and reopen the file since we already have a reader attached to it
  file.Close()
  file, err = os.Open("rhyme.txt")

  // Create a Scanner
  scanner := bufio.NewScanner(file)
  // Loop through the file, grabbing each line using .Scan() and .Text()
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}
