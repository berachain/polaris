// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"fmt"
)

// Color constants for terminal output.
const (
	resetColor   = "\033[0m"
	redColor     = "\033[31m"
	greenColor   = "\033[32m"
	yellowColor  = "\033[33m"
	blueColor    = "\033[34m"
	magentaColor = "\033[35m"
	cyanColor    = "\033[36m"
)

// LogRed prints a log message with red color.
func LogRed(msg string, args ...interface{}) {
	log(redColor, msg, args...)
}

// LogGreen prints a log message with green color.
func LogGreen(msg string, args ...interface{}) {
	log(greenColor, msg, args...)
}

// LogYellow prints a log message with yellow color.
func LogYellow(msg string, args ...interface{}) {
	log(yellowColor, msg, args...)
}

// LogBlue prints a log message with blue color.
func LogBlue(msg string, args ...interface{}) {
	log(blueColor, msg, args...)
}

// LogMagenta prints a log message with magenta color.
func LogMagenta(msg string, args ...interface{}) {
	log(magentaColor, msg, args...)
}

// LogCyan prints a log message with cyan color.
func LogCyan(msg string, args ...interface{}) {
	log(cyanColor, msg, args...)
}

// log is a helper function that prints a log message and key-value pairs
// with a specified color.
// colorCode: ANSI escape code for the desired color
// msg: The log message to be printed
// args: Key-value pairs to be printed
func log(colorCode string, msg string, args ...interface{}) {
	fmt.Printf("%s%s%s\n", colorCode, msg, resetColor)
	for i := 0; i < len(args); i += 2 {
		key := args[i]
		value := args[i+1]
		fmt.Printf("%s  %v: %v%s\n", colorCode, key, value, resetColor)
	}
}
