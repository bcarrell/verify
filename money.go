package verify

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	creditCardRegexp string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|" +
		"6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])" +
		"[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
)

func (v *verifier) CreditCard() *verifier {
	v.Results["isCreditCard"] = true
	r := regexp.MustCompile("[^0-9]+")
	card := r.ReplaceAllLiteralString(v.Query, "")
	r = regexp.MustCompile(creditCardRegexp)
	if !r.MatchString(card) {
		fmt.Println(card)
		v.Results["isCreditCard"] = false
		return v
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
	if sum%10 != 0 {
		v.Results["isCreditCard"] = false
	}
	return v
}
