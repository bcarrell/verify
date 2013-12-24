# verify

Super simple string validation for Go, inspired by node-validator.

    go get github.com/bcarrell/verify

## Example
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
		// if everything passed, tack on an .IsVerified() method call at the end
		isScientistVerified := scientist.IsVerified()
		fmt.Println(isScientistVerified) // true

	}

## API

Initialize by calling `verify.Verify(<str>)` and then use any of the following.
Where it makes sense, all methods will have an inverse method, usually prefixed
with `Isnt` or `Doesnt`, depending on grammar.

	.Email()
	.IsntEmail()

	.Url()
	.IsntUrl()

	.CreditCard()
	.IsntCreditCard()

	.Length(<int>)
	.IsntLength(<int>)

	.Int()
	.IsntInt()

	.Is(<string>)
	.Isnt(<string>)

	.IsEmpty()
	.IsntEmpty()

	.Contains(<string>)
	.DoesntContain(<string>)

	.IsIn(<[]string>)
	.IsntIn(<[]string>)

	.MinLength(<int>) <-- inclusive
	.MaxLength(<int>) <-- inclusive

## Contributing

I welcome any and all contributions.  If you'd like to request a verification
method, be sure to create an Issue describing it.  If you'd like to contribute
a verification method, please include an inverse method and corresponding test
cases with your addition.  Thanks!