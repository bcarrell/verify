package verify

type verifier struct {
	Query   string
	Results map[string][]bool
}

// verification initializer
func Verify(s string) *verifier {
	v := &verifier{Query: s, Results: make(map[string][]bool)}
	return v
}

// helper for creating a new result key or appending to an existing result slice
func (v *verifier) addVerification(key string, result bool) *verifier {
	if _, ok := v.Results[key]; ok {
		v.Results[key] = append(v.Results[key], result)
	} else {
		v.Results[key] = []bool{result}
	}
	return v
}

// a convenient finalizing function which confirms that all previously chained
// methods passed their verifications
func (v *verifier) IsVerified() bool {
	for _, result := range v.Results {
		for _, isValid := range result {
			if !isValid {
				return false
			}
		}
	}
	return true
}
