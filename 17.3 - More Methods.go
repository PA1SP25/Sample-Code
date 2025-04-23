/*
 * Author      : Luke Sathrum
 * Description : Writing methods to help with a DayOfYear structure
 */

package main

import "fmt"

/*
 * A structure to represent a Day of a Year
 */
type DayOfYear struct {
	month int
	day   int
}

/*
 * A helper method for month checking
 */
func isValidMonth(month int) bool {
	return month >= 1 && month <= 12
}

/*
 * A helper method for day checking
 */
func isValidDay(day int) bool {
	return day >= 1 && day <= 31
}

/*
 * This Method will setup the month field of a DayOfYear variable
 */
func (doy *DayOfYear) setMonth(month int) {
	if !isValidMonth(month) {
		month = 1
	}
	doy.month = month
}

/*
 * This Method will setup the day field of a DayOfYear variable
 */
func (doy *DayOfYear) setDay(day int) {
  if !isValidDay(day) {
		day = 1
	}
	doy.day = day
}

/*
 * This Method will setup a DayOfYear variable using the
 * supplied parameters
 */
func (doy *DayOfYear) setup(month, day int) {
  // Call methods to help
  doy.setMonth(month)
  doy.setDay(day)
}

/*
 * This method will output a DayOfYear variable
 * in a nice, pretty way
 */
func (doy *DayOfYear) output() {
	// Setup a map to translate integers -> strings
	monthTranslate := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	// Output a formattted date
	fmt.Println(monthTranslate[doy.month], doy.day)
}

func main() {
	// Setup a DayOfYear Strucutre Variable
	date := DayOfYear{11, 9}

	// Output our DayOfYear Variable
	date.output()

	// Change all the values in our DayOfYear
	date.setup(100, -8)

	// Output our DayOfYear Variable
	date.output()

  // Change the month
  date.setMonth(5)
	
  // Output our DayOfYear Variable
	date.output()

  // Change the month
  date.setDay(4)
	
  // Output our DayOfYear Variable
	date.output()
}
