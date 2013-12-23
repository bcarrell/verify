package verify

type verifier struct {
	Query   string
	Results map[string]bool
}

// verification initializer
func Verify(s string) *verifier {
	v := &verifier{Query: s, Results: make(map[string]bool)}
	return v
}

// a convenient finalizing function which confirms that all previously chained
// methods passed their verifications
func (v *verifier) IsVerified() bool {
	for _, isValid := range v.Results {
		if !isValid {
			return false
		}
	}
	return true
}
