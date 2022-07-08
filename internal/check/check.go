// Package that has several functions to perform checks on specific aspects such as errors, etc.
package check

import (
	"fmt"
	"log"
)

var (
	// allowedErrors private global variable of allowed errors
	allowedErrors = []string{"bucket already exists"}
)

// Err is a function that receives an error type variable and checks if it is
// different from null; if it is, it executes a panic.
func Err(err error) {
	if err != nil {
		panic(err)
	}
}

// ErrWithMsg is a function that receives a variable of type error and another
// of type string, checks if the error is different from null; if it is, it
// executes a panic and displays a message.
func ErrWithMsg(err error, msg string) {
	if err != nil {
		log.Fatalf("%v: %v", msg, err)
		panic(err)
	}
}

// CompareErr is a function that receives an error type variable and compares it
// with a global list of the check package containing allowed errors; if it is
// different from the allowed errors, a panic is executed.
func CompareErr(err error) {
	if err != nil {
		for _, v := range allowedErrors {
			if err.Error() != v {
				fmt.Println(err)
				panic(err)
			}
		}
	}
}
