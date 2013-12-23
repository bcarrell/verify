/**
 * This file implements validators one might find useful for web development.
 */
package verify

import "regexp"

const (
	emailRegexp string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9]" +
		"(?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	urlRegexp string = "^" +
		// protocol identifier
		"(?:(?:https?|ftp)://)" +
		// user:pass authentication
		"(?:\\S+(?::\\S*)?@)?" +
		// host name
		"(?:(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)" +
		// domain name
		"(?:\\.(?:[a-z\u00a1-\uffff0-9]+-?)*[a-z\u00a1-\uffff0-9]+)*" +
		// TLD identifier
		"(?:\\.(?:[a-z\u00a1-\uffff]{2,}))" +
		// port number
		"(?::\\d{2,5})?" +
		// resource path
		"(?:/[^\\s]*)?" +
		"$"
)

// verifies an email
// as with all regex-based email validations, this may return inaccurate results
func (v *verifier) Email() *verifier {
	result := regexp.MustCompile(emailRegexp).MatchString(v.Query)
	return v.addVerification("Email", result)
}

// verifies a url
func (v *verifier) Url() *verifier {
	result := regexp.MustCompile(urlRegexp).MatchString(v.Query)
	return v.addVerification("Url", result)
}
