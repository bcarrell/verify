package verify

import (
	"reflect"
	"testing"
)

// --- Helpers ---
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf(
			"Expected %v (type %v) - Got %v (type %v)",
			b,
			reflect.TypeOf(b),
			a,
			reflect.TypeOf(a),
		)
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf(
			"Did not expect %v (type %v) - Got %v (type %v)",
			b,
			reflect.TypeOf(b),
			a,
			reflect.TypeOf(a),
		)
	}
}

// --- End helpers ---

const (
	validEmail   string = "success@success.com"
	invalidEmail string = "fail@"
)

func Test_Verify(t *testing.T) {
	v := Verify(validEmail)
	refute(t, v, nil)
	expect(t, v.Query, validEmail)
}

func Test_IsValidEmail(t *testing.T) {
	v := Verify(validEmail).IsValidEmail()
	expect(t, v.Results["isEmail"], true)
	v = Verify(invalidEmail).IsValidEmail()
	expect(t, v.Results["isEmail"], false)
}
