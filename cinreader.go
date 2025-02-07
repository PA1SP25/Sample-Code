// Copyright 2019 J Boyd Trolinger. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package cinreader implements a robust handler for a variety
// of types for inputs from os.Stdin.

// TODO(jboydt): needs additional testing
// TODO(jboydt): consider adding additional types?
// TODO(jboydt): consider allowing clients to change error messages?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MessageType encodes the error messages CinReader may emit.
type MessageType int

// Defined values for MessageType.
const (
	emptyStringNotAllowed = iota
	invalidCharacter
	invalidCharacterSet
	invalidFloat
	invalidInteger
	invalidIntegerRange
	invalidStringSet
)

var default_messages = map[MessageType]string{
	emptyStringNotAllowed: "Empty input not allowed. Please re-enter: ",
	invalidCharacter:      "Not valid character input. Please re-enter: ",
	invalidCharacterSet:   "Input not in [%s]. Please re-enter: ",
	invalidFloat:          "Not a valid float. Please re-enter: ",
	invalidInteger:        "Not a valid integer. Please re-enter: ",
	invalidIntegerRange:   "Integer must be between %d - %d. Please re-enter: ",
	invalidStringSet:      "\"%s\" not valid input. Please re-enter: ",
}

// CinReader provides robust input handling from os.Stdin.
type CinReader struct {
	reader   *bufio.Reader
	messages map[MessageType]string
}

// NewCinReader initializes a CinReader.
func NewCinReader() *CinReader {
	return &CinReader{bufio.NewReader(os.Stdin), default_messages}
}

// ReadCharacter returns a rune/character from os.Stdin.
func (cin *CinReader) ReadCharacter() rune {
	var charInput rune
	var err error
  var countCharacters int
	for {
		charInput, _, err = cin.reader.ReadRune()
    // Added to handle multiple characters on input
    countCharacters = cin.reader.Buffered()
		cin.flush()
		if err != nil || countCharacters > 1 {
			fmt.Printf(cin.messages[invalidCharacter])
			continue
		} else if charInput == '\n' {
      fmt.Printf(cin.messages[emptyStringNotAllowed])
      continue
    }
		break
	}
	return charInput
}

// ReadCharacterSet returns a rune/character in charset
// from os.Stdin.
func (cin *CinReader) ReadCharacterSet(charset []rune) rune {
	var charInput rune
	var err error
  var countCharacters int
	for {
		charInput, _, err = cin.reader.ReadRune()
    // Added to handle multiple characters on input
    countCharacters = cin.reader.Buffered()
		cin.flush()
		if err != nil || countCharacters > 1 {
			fmt.Printf(cin.messages[invalidCharacter])
			continue
		} else {
			var inCharset bool
			for _, char := range charset {
				if charInput == char {
					inCharset = true
					break
				}
			}
			if !inCharset {
				fmt.Printf(cin.messages[invalidCharacterSet], string(charset))
				continue
			}
		}
		break
	}
	return charInput
}

// ReadFloat returns a float64 from os.Stdin.
func (cin *CinReader) ReadFloat() float64 {
	var floatInput float64
	var err error
	for {
		stringInput := cin.ReadString(false)
		floatInput, err = strconv.ParseFloat(stringInput, 64)
		if err != nil {
			fmt.Printf(cin.messages[invalidFloat])
			continue
		}
		break
	}
	return floatInput
}

// ReadInteger returns an int from os.Stdin.
func (cin *CinReader) ReadInteger() int {
	var intInput int
	var err error
	for {
		stringInput := cin.ReadString(false)
		intInput, err = strconv.Atoi(stringInput)
		if err != nil {
			fmt.Printf(cin.messages[invalidInteger])
			continue
		}
		break
	}
	return intInput
}

// ReadIntegerRange returns an int between min and max (inclusive)
// from os.Stdin.
func (cin *CinReader) ReadIntegerRange(min, max int) int {
	var intInput int
	var err error
	for {
		stringInput := cin.ReadString(false)
		intInput, err = strconv.Atoi(stringInput)
		if err != nil {
			fmt.Printf(cin.messages[invalidInteger])
			continue
		} else if intInput < min || intInput > max {
			fmt.Printf(cin.messages[invalidIntegerRange], min, max)
			continue
		}
		break
	}
	return intInput
}

// ReadString returns a string from os.Stdin.
func (cin *CinReader) ReadString(allowEmpty bool) string {

	var input string
	var err error
	for {
		input, err = cin.reader.ReadString('\n')
		input = strings.Trim(input, " \r\n")
		if err == nil {
			if len(input) == 0 && !allowEmpty {
				fmt.Printf(cin.messages[emptyStringNotAllowed])
				continue
			}
			break
		}
	}
	return input
}

// ReadStringSet returns a string in stringset from os.Stdin.
func (cin *CinReader) ReadStringSet(stringset []string, caseSensitive bool) string {
	if !caseSensitive {
		for i, str := range stringset {
			stringset[i] = strings.ToUpper(str)
		}
	}
	var strInput string
	for {
		strInput = cin.ReadString(false)
		if !caseSensitive {
			strInput = strings.ToUpper(strInput)
		}
		var validString bool
		for _, str := range stringset {
			if str == strInput {
				validString = true
				break
			}
		}
		if !validString {
			fmt.Printf(cin.messages[invalidStringSet], strInput)
			continue
		}
		break
	}
	return strInput
}

// flush the input buffer.
func (cin *CinReader) flush() {
	cin.reader.Discard(cin.reader.Buffered())
}
