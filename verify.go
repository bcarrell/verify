package verify

type verifier struct {
	Query   string
	Results map[string]bool
	IsValid bool
}

func Verify(s string) *verifier {
	v := &verifier{Query: s, Results: make(map[string]bool)}
	return v
}
