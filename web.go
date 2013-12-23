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

// checks for a valid email
// as with all regex-based email validations, this may return inaccurate results
func (v *verifier) Email() *verifier {
	r := regexp.MustCompile(emailRegexp)
	v.Results["Email"] = r.MatchString(v.Query)
	return v
}

func (v *verifier) Url() *verifier {
	r := regexp.MustCompile(urlRegexp)
	v.Results["Url"] = r.MatchString(v.Query)
	return v
}
