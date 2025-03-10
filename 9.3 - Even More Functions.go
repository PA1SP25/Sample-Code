/*
 * Author      : Luke Sathrum
 * Description : Even More Function Examples
 */

package main

import "fmt"

// totalCost computes the total cost, including a 5% sales tax
// on numberParam items at a cost of priceParam each.
func totalCost(numberParam int, priceParam float64) float64 {
	// Set up sales tax
	const kTaxRate float64 = 0.05

	// declare subtotal variable
	var subtotal float64

	// Output some information about the function
	fmt.Println("We are in function totalCost")
	fmt.Println("The parameter numberParam has a value of", numberParam)
	fmt.Println("The parameter priceParam has a value of", priceParam)

	// Calculate the Subtotal
	subtotal = priceParam * float64(numberParam)
	return subtotal + (subtotal * kTaxRate)
}

func main() {
	// Set up our Variables
	var (
		price  float64
		bill   float64
		number int
	)

	// Get price and number from the user
	fmt.Print("Enter the number of items purchased: ")
	fmt.Scan(&number)
	fmt.Print("Enter the price per item $")
	fmt.Scan(&price)

	// Call function totalCost
	bill = totalCost(number, price)

	// Output the results
	fmt.Printf("%d items @ $%.2f each.\n", number, price)
	fmt.Printf("Final bill (including tax): $%.2f\n", bill)
}