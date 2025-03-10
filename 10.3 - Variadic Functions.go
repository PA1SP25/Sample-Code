/*
 * Author      : Luke Sathrum
 * Description : Examples of Variadic Functions
 */

package main

import "fmt"

/*
 * This function will add all of the arguments together
 * and return the sum
 */
func sum(vals ...int) int {
  sum := 0
  for _, val := range vals {
    sum += val
  }
  return sum
}


func main() {
  // call sum with varying number of arguments
  fmt.Println("sum()              ->", sum())
  fmt.Println("sum(1)             ->", sum(1))
  fmt.Println("sum(1, 2)          ->", sum(1, 2))
  fmt.Println("sum(1, 2, 3)       ->", sum(1, 2, 3))
  fmt.Println("sum(1, 2, 3, 4)    ->", sum(1, 2, 3, 4))
  fmt.Println("sum(1, 2, 3, 4, 5) ->", sum(1, 2, 3, 4, 5))
}