/*
 * Author      : Luke Sathrum
 * Description : Examples of Functions
 */

package main

import "fmt"

// outputs Hello to the Console
// A simple function that does not have any
// parameters and doesn't produce anything
func hello() {
	fmt.Println("Hello")
}

// outputs a string to the Console
// The string comes in via a parameter
func output(word string) {
	fmt.Println(word)
}

// Returns the meaning of life
func meaningOfLife() int {
	return 42
}

// adds (sums) the values of the 2 parameters
// Produces (returns) the sum
func add(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	// Calling a function
	hello()

	// Calling a function w/ a parameter
	// by passing in an argument (value)
	// Argument can be hard coded
	output("Hello World!")
	// Argument can be a variable
	val := "a Variable!"
	output(val)
	// Argument can be an expression
	output("An expression with " + val)

	// Calling a function w/ a return value
	// Need to do something with the value
	// meaningOfLife() <- Does Nothing
	fmt.Println(meaningOfLife())

	// Calling fuction that takes input (via parameter)
	// AND produces a result
	sum := add(1, 2)
	fmt.Println(sum)

	// You can use function calls anywhere you use variables
	sum = add(meaningOfLife(), 100)
	fmt.Println(sum)
}
