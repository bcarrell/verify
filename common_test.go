package verify

import "testing"

var (
	threeCharacterNames = []string{
		"bob",
		"mat",
		"jim",
		"sue",
	}
	validIntegers = []string{
		"123",
		"4",
		"9993",
	}
	invalidIntegers = []string{
		"12 3",
		" 123",
		" 1",
		"1,200",
		"fm",
		"dskq",
		" a",
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
		v = Verify(name).IsntLength(123).IsVerified()
		expect(t, v, true)
	}
}

func Test_Integers(t *testing.T) {
	for _, integer := range validIntegers {
		v := Verify(integer).Int().IsVerified()
		expect(t, v, true)
		v = Verify(integer).IsntInt().IsVerified()
		refute(t, v, true)
	}
	for _, integer := range invalidIntegers {
		v := Verify(integer).Int().IsVerified()
		expect(t, v, false)
		v = Verify(integer).IsntInt().IsVerified()
		refute(t, v, false)
	}
}

func Test_Is(t *testing.T) {
	v := Verify("seafood").Is("seafood").IsVerified()
	expect(t, v, true)
	v = Verify("seafood").Is("Seafood").IsVerified()
	expect(t, v, false)
	v = Verify("seafood").Is("Poultry").IsVerified()
	expect(t, v, false)
}

func Test_Isnt(t *testing.T) {
	v := Verify("seafood").Isnt("poultry").IsVerified()
	expect(t, v, true)
	v = Verify("seafood").Isnt("seafood").IsVerified()
	expect(t, v, false)
}

func Test_IsntEmpty(t *testing.T) {
	v := Verify("seafood").IsntEmpty().IsVerified()
	expect(t, v, true)
	v = Verify("").IsntEmpty().IsVerified()
	expect(t, v, false)
}

func Test_IsEmpty(t *testing.T) {
	v := Verify("").IsEmpty().IsVerified()
	expect(t, v, true)
	v = Verify("seafood").IsEmpty().IsVerified()
	expect(t, v, false)
}

func Test_Contains(t *testing.T) {
	v := Verify("team").Contains("i").IsVerified()
	expect(t, v, false)
	v = Verify("team").Contains("e").Contains("a").IsVerified()
	expect(t, v, true)
	v = Verify("team").Length(4).Contains("e").IsVerified()
	expect(t, v, true)
	v = Verify("team").MaxLength(20).MinLength(5).Contains("e").IsVerified()
	expect(t, v, false)
	v = Verify("team").MaxLength(8).MinLength(2).IContains("E").IsVerified()
	expect(t, v, true)
}

func Test_DoesntContain(t *testing.T) {
	v := Verify("team").DoesntContain("i").IsVerified()
	expect(t, v, true)
	v = Verify("team").DoesntContain("e").IsVerified()
	expect(t, v, false)
	v = Verify("team").MaxLength(8).IDoesntContain("yy").IsVerified()
	expect(t, v, true)
}
