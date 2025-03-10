/*
 * Author      : Luke Sathrum
 * Description : Function Examples of Decomposition and Abstraction
 */

package main

import "fmt"

// An implementation of Summing to a Value
// Takes an integer to sum the values from 1 - integer
// Returns the total
func sumTo(val int) int {
	return (val / 2.0) * (val + 1)
}

// A Basic Implementation of Exponents
func raiseToPower(base, exponent int) int {
	solution := 1
	for i := 0; i < exponent; i++ {
		solution *= base
	}
	return solution
}

func main() {
	// Call sumTo and Output the Results
	fmt.Println(sumTo(200))

	// Make it interactive. We get input into function
	// via Parameters, but the arguments (values) can
	// come from anywhere
	var base, exponent int
	fmt.Print("Give me the base of an exponent, followed by the exponent: ")
	fmt.Scan(&base, &exponent)
	fmt.Printf("%d^%d = %d\n", base, exponent, raiseToPower(base, exponent))
}
