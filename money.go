// Implements verifications for handling financial information.
package verify

import (
	"regexp"
	"strconv"
)

const (
	creditCardRegexp string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|" +
		"6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])" +
		"[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
)

// verifies a credit card number
// we consider a credit card to be valid if it passes the constant
// creditCardRegexp and passes the Luhn algorithm
// This method should currently verify the following card formats:
//     - Visa
//     - MasterCard
//     - AmEx
//     - Discover
//     - Diners Club
//     - JCB
func (v *verifier) CreditCard() *verifier {
	card := regexp.MustCompile("[^0-9]+").ReplaceAllLiteralString(v.Query, "")
	r := regexp.MustCompile(creditCardRegexp)
	if !r.MatchString(card) {
		return v.addVerification("CreditCard", false)
	}

	// Luhn algorithm
	// http://en.wikipedia.org/wiki/Luhn_algorithm
	sum := 0
	shouldDouble := false
	for i := len(card) - 1; i >= 0; i-- {
		toInt, _ := strconv.Atoi(string(card[i]))
		if shouldDouble {
			doubled := toInt * 2
			if doubled >= 10 {
				sum += (doubled % 10) + 1
			} else {
				sum += doubled
			}
		} else {
			sum += toInt
		}
		shouldDouble = !shouldDouble
	}
	return v.addVerification("CreditCard", sum%10 == 0)
}
