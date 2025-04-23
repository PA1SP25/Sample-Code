/*
 * Author      : Luke Sathrum
 * Description : Writing functions to help with a DayOfYear structure
 */

package main

import "fmt"

type DayOfYear struct {
  month int
  day int
}

/*
 * This function will setup a DayOfYear variable using the
 * supplied parameters
 */
func doySetup(month, day int, doy *DayOfYear) {
  doy.month = month
  doy.day = day
}

/*
 * This function will output a DayOfYear variable
 * in a nice, pretty way
 */
func doyOutput(doy DayOfYear) {
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
  doyOutput(date)
  
  // Change the values in our DayOfYear
  doySetup(11, 8, &date)

  // Output our DayOfYear Variable
  doyOutput(date)
}