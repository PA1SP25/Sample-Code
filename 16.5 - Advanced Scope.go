/*
 * Author      : Luke Sathrum
 * Description : Class Examples of Block Scope and for loop scope
 */

package main

import (
	"fmt"
)

// Program Starts Here
func main() {
	// We have blocks inside of blocks
	// Scope is loacl to the block
	var number int = 1
	{
		var number int = 2
		fmt.Println("Inside the inner block")
		fmt.Println("The number is:", number)
	}
	fmt.Println("Inside the outer block")
	fmt.Println("The number is:", number)

	// for loop example
	// i := 0
	for i := 1; i < 5; i++ {
		fmt.Println("The value of i is:", i)
	}
	// What if we want to access i here?
	// fmt.Println("The value of i is:", i)
}
