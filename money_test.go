package verify

import "testing"

var (
	validCreditCardNumbers = []string{
		"378282246310005",
		"371449635398431",
		"378734493671000",
		"30569309025904",
		"38520000023237",
		"6011111111111117",
		"6011000990139424",
		"3530111333300000",
		"3566002020360505",
		"5555555555554444",
		"5105105105105100",
		"4111111111111111",
		"4012888888881881",
		"4222222222222",
	}
	invalidCreditCardNumbers = []string{
		"1234",
		"123124923j",
		"asdfasdf32",
		"234jdm",
	}
)

// test for .IsCreditCard()
func Test_IsCreditCard(t *testing.T) {
	for _, card := range validCreditCardNumbers {
		v := Verify(card).IsCreditCard().IsVerified()
		expect(t, v, true)
	}

	for _, card := range invalidCreditCardNumbers {
		v := Verify(card).IsCreditCard().IsVerified()
		refute(t, v, true)
	}
}

// test for .IsntCreditCard()
func Test_IsntCreditCard(t *testing.T) {
	for _, card := range invalidCreditCardNumbers {
		v := Verify(card).IsntCreditCard().IsVerified()
		expect(t, v, true)
	}

	for _, card := range validCreditCardNumbers {
		v := Verify(card).IsntCreditCard().IsVerified()
		refute(t, v, true)
	}
}
