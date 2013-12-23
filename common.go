// Implements general use case verifications.
package verify

import (
	"strconv"
	"strings"
)

// verifies a string has a length equal to `length`
func (v *verifier) Length(length int) *verifier {
	return v.addVerification("Length", len(v.Query) == length)
}

// verifies a string has a length greater than or equal to `length`
func (v *verifier) MinLength(length int) *verifier {
	return v.addVerification("MinLength", len(v.Query) >= length)
}

// verifies a string has a length less than or equal to `length`
func (v *verifier) MaxLength(length int) *verifier {
	return v.addVerification("MaxLength", len(v.Query) <= length)
}

// verifies a string is a string representation of a valid integer
func (v *verifier) Int() *verifier {
	_, err := strconv.Atoi(v.Query)
	return v.addVerification("Int", err == nil)
}

// verifies a string is equivalent to another string
func (v *verifier) Is(comparison string) *verifier {
	return v.addVerification("Is", v.Query == comparison)
}

// verifies a string is not equivalent to another string
func (v *verifier) Isnt(comparison string) *verifier {
	return v.addVerification("Isnt", v.Query != comparison)
}

// verifies a string is zero-length (aka not empty)
func (v *verifier) IsEmpty() *verifier {
	return v.addVerification("IsEmpty", len(v.Query) == 0)
}

// verifies a string is not zero-length (aka not empty)
func (v *verifier) IsntEmpty() *verifier {
	return v.addVerification("IsntEmpty", len(v.Query) > 0)
}

// verifies a string contains an `inner` string
func (v *verifier) Contains(inner string) *verifier {
	return v.addVerification("Contains", strings.Contains(v.Query, inner))
}

// verifies a string does not contain an `inner` string
func (v *verifier) DoesntContain(inner string) *verifier {
	return v.addVerification("DoesntContain", !strings.Contains(v.Query, inner))
}
