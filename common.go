package verify

func (v *verifier) MinLength(length int) *verifier {
	v.Results["isMinLength"] = len(v.Query) >= length
	return v
}

func (v *verifier) MaxLength(length int) *verifier {
	v.Results["isMaxLength"] = len(v.Query) <= length
	return v
}

func (v *verifier) Length(length int) *verifier {
	v.Results["isLength"] = len(v.Query) == length
	return v
}
