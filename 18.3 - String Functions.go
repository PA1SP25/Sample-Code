/*
 * Author      : Luke Sathrum
 * Description : Examples of string functions
 */

package main

import (
	"fmt"
	s "strings"
)

func main() {
	// Contains
  fmt.Println("---Contains()---")
	str := "Hello World!"
  str2 := "Wor?"
  fmt.Printf("Does %q contain %q? %t\n\n", str, str2, s.Contains(str, str2))
	

  // Count
  fmt.Println("---Count()---")
	str = "Foo Bar FooBar"
  str2 = "o"
  fmt.Printf("How many %qs in %q?   %d\n", str, str2, s.Count(str, str2))
  str2 = "oo"
  fmt.Printf("How many %qs in %q?   %d\n\n", str, str2, s.Count(str, str2))
  
  // Has Prefix
  fmt.Println("---HasPrefix()---")
	str = "Hello World!"
  str2 = "He"
  fmt.Printf("Does %q start with %q?  %t\n", str, str2, s.HasPrefix(str, str2))
  str2 = "hel"
  fmt.Printf("Does %q start with %q? %t\n\n", str, str2, s.HasPrefix(str, str2))
  
  // Has Prefix
  fmt.Println("---HasSuffix()---")
	str = "Hello World!"
  str2 = "d!"
  fmt.Printf("Does %q end with %q? %t\n", str, str2, s.HasSuffix(str, str2))
  str2 = "He"
  fmt.Printf("Does %q end with %q? %t\n\n", str, str2, s.HasSuffix(str, str2))

  // Index
  fmt.Println("---Index()---")
	str = "Foo Bar FooBar"
  str2 = "ar"
  fmt.Printf("%q is found in %q starting in location %d\n", str2, str, s.Index(str, str2))
  str2 = "ba"
  fmt.Printf("%q is found in %q starting in location %d\n\n", str2, str, s.Index(str, str2))
  
  // Join
  fmt.Println("---Join()---")
	strSlice := []string{"H", "E", "L", "L", "O"}
  str = s.Join(strSlice, "")
  fmt.Printf("s.Join(%#v, \"\")   becomes: %q\n", strSlice, str)
  str = s.Join(strSlice, ", ")
  fmt.Printf("s.Join(%#v, \", \") becomes: %q\n\n", strSlice, str)

  // Repeat
  fmt.Println("---Repeat()---")
  str = "foo"
  fmt.Printf("s.Repeat(%q, 5) becomes: %q\n", str, s.Repeat(str, 5))
  str = "."
  fmt.Printf("s.Repeat(%q, 5) becomes:   %q\n\n", str, s.Repeat(str, 5))
	
  // Replace
  fmt.Println("---Replace()---")
  str = "Foo Foo Foo"
  str2 = "o"
  str3 := "0"
  fmt.Printf("We can replace the first 3 %qs with %q in the string %q: %q\n", str2, str3, str, s.Replace(str, str2, str3, 3))
  str2 = "Foo"
  str3 = "Foobar"
  fmt.Printf("We can replace ALL %qs with %q in the string %q:  %q\n\n", str2, str3, str, s.Replace(str, str2, str3, -1))
  
  // Split
  fmt.Println("---Split()---")
	str = "A, B, C, D, E, F"
  strSlice = s.Split(str, "")
  fmt.Printf("s.Split(%q, \"\") becomes\n%#v\n", str, strSlice)
  strSlice = s.Split(str, ", ")
  fmt.Printf("s.Split(%q, \", \") becomes\n%#v\n\n", str, strSlice)

  // ToLower
  fmt.Println("---ToLower()---")
	str = "PLEASE STOP YELLING!"
  fmt.Printf("s.ToLower(%q) becomes %q\n", str, s.ToLower(str))
  str = "Other Characters are OK 123 %^&"
  fmt.Printf("s.ToLower(%q) becomes %q\n\n", str, s.ToLower(str))

  // ToUpper
  fmt.Println("---ToUpper()---")
	str = "Please be louder!"
  fmt.Printf("s.ToUpper(%q) becomes %q\n\n", str, s.ToUpper(str))
}
