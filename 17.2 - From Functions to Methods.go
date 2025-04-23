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
  day int
}

/*
 * This Method will setup a DayOfYear variable using the
 * supplied parameters
 */
func (doy *DayOfYear) setup(month, day int) {
  // This is a good place to do some error handling as well
  if month < 1 || month > 12 {
    month = 1
  }
  doy.month = month
  if day < 1 || day > 31 {
    day = 1
  }
  doy.day = day
}

/*
 * This method will output a DayOfYear variable
 * in a nice, pretty way
 */
func (doy *DayOfYear) output() {
  // Setup a map to translate integers -> strings
  monthTranslate := map[int]string{
     1: "January",
     2: "February",
     3: "March",
     4: "April",
     5: "May",
     6: "June",
     7: "July",
     8: "August",
     9: "September",
    10: "October",
    11: "November",
    12: "December",
  }
  // Output a formattted date
  fmt.Println(monthTranslate[doy.month], doy.day)
}

func main() {
  // Setup a DayOfYear Strucutre Variable
  date := DayOfYear{1, 15}
  
  // Output our DayOfYear Variable
  date.output()
  
  // Change the values in our DayOfYear
  date.setup(11, 8)

  // Output our DayOfYear Variable
  date.output()
}
