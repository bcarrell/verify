# verify

String validation utilities for Golang.

### Installation
    go get github.com/bcarrell/verify

# Example
	package main

	import (
		"fmt"

		"github.com/bcarrell/verify"
	)

	func main() {
		// basic usage -- chaining is encouraged, and `verify` is meant to have
		// an API approaching plain english (ie. DoesntContain(..) instead of
		// NotContain(..))
		scientist := verify.Verify("Gaius Baltar").
			MinLength(8).
			MaxLength(200).
			Contains("Gaius").
			Contains("Baltar").
			DoesntContain("cylon")

		// You can inspect each verification result independently.
		// Each method stores a string key of []bool values within a Results map.
		// The key is identical to its function name.
		// Each []bool item result is in the order of the functions called.
		fmt.Println(scientist.Results["MinLength"])     // [true]
		fmt.Println(scientist.Results["MaxLength"])     // [true]
		fmt.Println(scientist.Results["Contains"])      // [true true]
		fmt.Println(scientist.Results["DoesntContain"]) // [true]

		// or, if you don't care about individual results and just want to know
		// if everything passed, tack on an .isVerified() method call at the end
		isScientistVerified := scientist.IsVerified()
		fmt.Println(isScientistVerified) // true

	}
