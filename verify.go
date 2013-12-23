package verify

type verifier struct {
	Query   string
	Results map[string]bool
}

func Verify(s string) *verifier {
	v := &verifier{Query: s, Results: make(map[string]bool)}
	return v
}

func (v *verifier) IsVerified() bool {
	for _, isValid := range v.Results {
		if !isValid {
			return false
		}
	}
	return true
}
