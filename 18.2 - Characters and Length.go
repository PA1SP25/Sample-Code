/*
 * Author      : Luke Sathrum
 * Description : Examples of string manipulation
 */

package main

import (
  "fmt"
  s "strings" // Alias the packages as s
)

func main() {
	// Create our string
  fmt.Println("---The Initial String---")
  str := "HELLO"
  fmt.Println(str)

  fmt.Println("\n---Getting Single Characters---")

  // Getting Characters (as bytes)
  fmt.Println(str[0])

  // Going out of range crashes the program
  // fmt.Println(str[10])

  // Create a byte variable
  var char byte
  char = str[3]
  fmt.Println(char)

  // Ways of outputting as the actual character
  fmt.Println(string(char))
  // %T gives us the type
  fmt.Printf("%T\n", char)
  // %c gives us the "Code Point, i.e. Rune"
  fmt.Printf("%c\n", char)
  // %U givs us the unicode format
  fmt.Printf("%U\n", char)
  // %#U gives us the unicode and Code Point
  fmt.Printf("%#U\n", char)

  // Getting multiple characters (called a substring)
  fmt.Println("\n---Getting Multiple Characters---")
  str2 := str[1:2]
  fmt.Println(str2)
  str2 = str[1:]
  fmt.Println(str2)
  str2 = str[:2]
  fmt.Println(str2)
  str2 = str[:]
  fmt.Println(str2)


  // Length
  fmt.Println("\n---Working with Length len()---")
  fmt.Println("The length of our string is:", len(str))

  // Some loops
  // For Range (becomes a rune)
  for _, c := range str {
    fmt.Printf("%c(%v) is of type %T\n", c, c, c)
  }

  // Regular for (becomes a byte)
  for i := 0; i < len(str); i++ {
    fmt.Printf("%c(%v) is of type %T\n", str[i], str[i], str[i])
  }

  // For Loop with Fun
  for i := 0; i < len(str); i++ {
    fmt.Printf("%s%c\n", s.Repeat(" ", 2*i), str[i])
  }

}
