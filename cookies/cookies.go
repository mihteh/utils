package cookies

import (
	"net/http"
)

// HttpCookies is a slice of pointers to http.Cookie objects
type HttpCookies []*http.Cookie

// New parses string containing cookies and returns a slice of pointers to http.Cookie objects
func New(cookie string) HttpCookies {
	header := http.Header{}
	header.Add("Cookie", cookie)
	request := http.Request{Header: header}
	return request.Cookies()
}

// Get returns a cookie with Name equals to key
func (cookies HttpCookies) Get(key string) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == key {
			return cookie
		}
	}
	return nil
}
