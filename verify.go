package verify

import "regexp"

type verifier struct {
	Query   string
	Results map[string]bool
}

func Verify(s string) *verifier {
	v := &verifier{Query: s, Results: make(map[string]bool)}
	return v
}

func (v *verifier) IsValidEmail() *verifier {
	r := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	v.Results["isEmail"] = r.MatchString(v.Query)
	return v
}
