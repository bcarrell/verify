package verify

import (
	"strconv"
	"strings"
)

func (v *verifier) MinLength(length int) *verifier {
	v.Results["MinLength"] = len(v.Query) >= length
	return v
}

func (v *verifier) MaxLength(length int) *verifier {
	v.Results["MaxLength"] = len(v.Query) <= length
	return v
}

func (v *verifier) Length(length int) *verifier {
	v.Results["Length"] = len(v.Query) == length
	return v
}

func (v *verifier) Int() *verifier {
	v.Results["Int"] = true
	_, err := strconv.Atoi(v.Query)
	if err != nil {
		v.Results["Int"] = false
	}
	return v
}

func (v *verifier) Is(comparison string) *verifier {
	v.Results["Is"] = v.Query == comparison
	return v
}

func (v *verifier) Isnt(comparison string) *verifier {
	v.Results["Isnt"] = v.Query != comparison
	return v
}

func (v *verifier) IsEmpty() *verifier {
	v.Results["IsEmpty"] = len(v.Query) == 0
	return v
}

func (v *verifier) IsntEmpty() *verifier {
	v.Results["IsntEmpty"] = len(v.Query) > 0
	return v
}

func (v *verifier) Contains(inner string) *verifier {
	v.Results["Contains"] = strings.Contains(v.Query, inner)
	return v
}

func (v *verifier) DoesntContain(inner string) *verifier {
	v.Results["DoesntContain"] = !strings.Contains(v.Query, inner)
	return v
}
