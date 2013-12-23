// Common is a collection of methods for more general string verifications
package verify

import (
	"strconv"
	"strings"
)

// verifies a string has a length greater than or equal to `length`
func (v *verifier) MinLength(length int) *verifier {
	v.Results["MinLength"] = len(v.Query) >= length
	return v
}

// verifies a string has a length less than or equal to `length`
func (v *verifier) MaxLength(length int) *verifier {
	v.Results["MaxLength"] = len(v.Query) <= length
	return v
}

// verifies a string has a length equal to `length`
func (v *verifier) Length(length int) *verifier {
	v.Results["Length"] = len(v.Query) == length
	return v
}

// verifies a string is a string representation of a valid integer
func (v *verifier) Int() *verifier {
	v.Results["Int"] = true
	_, err := strconv.Atoi(v.Query)
	if err != nil {
		v.Results["Int"] = false
	}
	return v
}

// verifies a string is equivalent to another string
func (v *verifier) Is(comparison string) *verifier {
	v.Results["Is"] = v.Query == comparison
	return v
}

// verifies a string is not equivalent to another string
func (v *verifier) Isnt(comparison string) *verifier {
	v.Results["Isnt"] = v.Query != comparison
	return v
}

// verifies a string is zero-length (aka not empty)
func (v *verifier) IsEmpty() *verifier {
	v.Results["IsEmpty"] = len(v.Query) == 0
	return v
}

// verifies a string is not zero-length (aka not empty)
func (v *verifier) IsntEmpty() *verifier {
	v.Results["IsntEmpty"] = len(v.Query) > 0
	return v
}

// verifies a string contains an `inner` string
func (v *verifier) Contains(inner string) *verifier {
	v.Results["Contains"] = strings.Contains(v.Query, inner)
	return v
}

// verifies a string does not contain an `inner` string
func (v *verifier) DoesntContain(inner string) *verifier {
	v.Results["DoesntContain"] = !strings.Contains(v.Query, inner)
	return v
}
