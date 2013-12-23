package verify

import "testing"

var (
	threeCharacterNames = []string{
		"bob",
		"mat",
		"jim",
		"sue",
	}
)

func Test_Length(t *testing.T) {
	for _, name := range threeCharacterNames {
		v := Verify(name).MinLength(2).IsVerified()
		expect(t, v, true)
		v = Verify(name).MinLength(3).IsVerified()
		expect(t, v, true)
		v = Verify(name).MinLength(3).MaxLength(8).IsVerified()
		expect(t, v, true)
		v = Verify(name).MinLength(3).MaxLength(3).IsVerified()
		expect(t, v, true)
		v = Verify(name).MinLength(4).MaxLength(8).IsVerified()
		expect(t, v, false)
		v = Verify(name).Length(3).IsVerified()
		expect(t, v, true)
		v = Verify(name).Length(2).IsVerified()
		expect(t, v, false)
		v = Verify(name).Length(4).IsVerified()
		expect(t, v, false)
	}

}
