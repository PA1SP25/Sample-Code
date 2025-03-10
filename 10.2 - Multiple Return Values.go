/*
 * Author      : Luke Sathrum
 * Description : Examples of Returning Multiple Values
 */

package main

import "fmt"

/*
 * A function that returns two values
 */
func greeting() (string, string) {
  return "Hello", "World!"
}


/*
 * Performs addition and subtraction on parameters and
 * returns the result of both operations
 */
func math(x, y int) (int, int) {
  return x + y, x - y
}

func main() {
  // Call greeting(), storing the two returned values
  word, word2 := greeting()
  fmt.Println(word, word2)

  // Call math(), storing the two returned values
  val1, val2 := math(8, 5)
  fmt.Printf("8 + 5 = %d, 8 - 5 = %d\n", val1, val2)
}
