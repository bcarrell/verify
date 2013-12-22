package verify

type verifier struct {
	String  string
	Results map[string]bool
}

func Verify(s string) *verifier {
	v := &verifier{s}
	return v
}

func (v *verifier) IsValidEmail() *verifier {

}
