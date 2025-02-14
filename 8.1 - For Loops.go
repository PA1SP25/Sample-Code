/*
 * Author      : Luke Sathrum
 * Description : Introduction to for loops
 */

package main

import "fmt"

func main() {
	// Create our variables
	var (
		countDown int
		iteration = 1
	)

	// Get input from the user
	fmt.Print("How many greetings do you want? ")
	fmt.Scan(&countDown)

	for countDown > 0 {
		// Display greeting and the current iteration to the user
		fmt.Printf("Hello! (Iteration #%d)\n", iteration)
		// Decrement our counting variable
		countDown--
		// Increment our iteration variable
		iteration++
	}
	fmt.Println("\nAll done with the loop")

}
