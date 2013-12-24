package verify

import "testing"

const (
	validEmail   string = "success@success.com"
	invalidEmail string = "fail@"
)

var (
	validUrls = []string{
		"http://foo.com/blah_blah",
		"http://foo.com/blah_blah/",
		"http://foo.com/blah_blah_(wikipedia)",
		"http://foo.com/blah_blah_(wikipedia)_(again)",
		"http://www.example.com/wpstyle/?p=364",
		"https://www.example.com/foo/?bar=baz&inga=42&quux",
		"http://✪df.ws/123",
		"http://userid:password@example.com:8080",
		"http://userid:password@example.com:8080/",
		"http://userid@example.com",
		"http://userid@example.com/",
		"http://userid@example.com:8080",
		"http://userid@example.com:8080/",
		"http://userid:password@example.com",
		"http://userid:password@example.com/",
		"http://➡.ws/䨹",
		"http://⌘.ws",
		"http://⌘.ws/",
		"http://foo.com/blah_(wikipedia)#cite-1",
		"http://foo.com/blah_(wikipedia)_blah#cite-1",
		"http://foo.com/unicode_(✪)_in_parens",
		"http://foo.com/(something)?after=parens",
		"http://☺.damowmow.com/",
		"http://code.google.com/events/#&product=browser",
		"http://j.mp",
		"ftp://foo.bar/baz",
		"http://foo.bar/?q=Test%20URL-encoded%20stuff",
		"http://مثال.إختبار",
		"http://例子.测试",
		"http://उदाहरण.परीक्षा",
		"http://-.~_!$&'()*+,;=:%40:80%2f::::::@example.com",
		"http://1337.net",
		"http://a.b-c.de",
	}
	invalidUrls = []string{
		"http://",
		"http://.",
		"http://..",
		"http://../",
		"http://?",
		"http://??",
		"http://??/",
		"http://#",
		"http://##",
		"http://##/",
		"http://foo.bar?q=Spaces should be encoded",
		"//",
		"//a",
		"///a",
		"///",
		"http:///a",
		"foo.com",
		"rdar://1234",
		"h://test",
		"http:// shouldfail.com",
		":// should fail",
		"http://foo.bar/foo(bar)baz quux",
		"ftps://foo.bar/",
		"http://-error-.invalid/",
		"http://a.b--c.de/",
		"http://-a.b.co",
		"http://a.b-.co",
		"http://0.0.0.0",
		"http://10.1.1.0",
		"http://10.1.1.255",
		"http://224.1.1.1",
		"http://1.1.1.1.1",
		"http://123.123.123",
		"http://3628126748",
		"http://.www.foo.bar/",
		"http://www.foo.bar./",
		"http://.www.foo.bar./",
		"http://10.1.1.1",
		"http://10.1.1.254",
	}
)

func Test_IsEmail(t *testing.T) {
	v := Verify(validEmail).Email().IsVerified()
	expect(t, v, true)
	v = Verify(invalidEmail).Email().IsVerified()
	expect(t, v, false)
}

func Test_IsntEmail(t *testing.T) {
	v := Verify(validEmail).IsntEmail().IsVerified()
	expect(t, v, false)
	v = Verify(invalidEmail).IsntEmail().IsVerified()
	expect(t, v, true)
}

func Test_IsUrl(t *testing.T) {
	for _, url := range invalidUrls {
		v := Verify(url).Url().IsVerified()
		expect(t, v, false)
	}

	for _, url := range validUrls {
		v := Verify(url).Url().IsVerified()
		expect(t, v, true)
	}
}

func Test_IsntUrl(t *testing.T) {
	for _, url := range validUrls {
		v := Verify(url).IsntUrl().IsVerified()
		expect(t, v, false)
	}

	for _, url := range invalidUrls {
		v := Verify(url).IsntUrl().IsVerified()
		expect(t, v, true)
	}
}
