package util

import (
	"net/url"
	"strings"
)

// ValidateAndParseURL := Convert string to a valid URL, throw error if parsing fails
func ValidateAndParseURL(urlString string) (*url.URL, error) {
	// Append
	u := AddPrefixHttp(urlString)
	// Parse
	return url.ParseRequestURI(u)
}

// Assumption, if URL does not contain http prefix, this code would add it before making a GET call.
// AddPrefixHttp: Add Http prefix if missing in the input
func AddPrefixHttp(urlString string) string {
	prefix := "http://"
	if urlString != "" && !strings.HasPrefix(urlString, prefix) {
		urlString = prefix + urlString
	}
	return urlString
}
