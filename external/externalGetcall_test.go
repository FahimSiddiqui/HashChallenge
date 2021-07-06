package external

import (
	"net/url"
	"strings"
	"testing"
)

func TestExternalGetCall(t *testing.T) {
	url := url.URL{Scheme: "http", Host: "www.google.com"}
	responseStruct, err := GetUrlResponse(&url)
	if err != nil || responseStruct.HashedOutput == "" || strings.Compare(responseStruct.InputUrl, url.String()) != 0 {
		t.Error()
	}
}
