package util

import (
	"strings"
	"testing"
)

func TestURLParsing(t *testing.T) {
	inputUrl := "http://www.google.com"
	parseUrl, err := ValidateAndParseURL(inputUrl)
	if err != nil || strings.Compare(inputUrl, parseUrl.String()) != 0 {
		t.Fail()
	}

	inputUrl = "abc"
	parseUrl, err = ValidateAndParseURL(inputUrl)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	// Empty URL
	inputUrl = ""
	parseUrl, err = ValidateAndParseURL(inputUrl)
	// Error should occur, in order to make it a success, checking nil.
	if err == nil {
		t.Log(err)
		t.Fail()
	}
}

func TestAddingPrefix(t *testing.T) {
	inputUrl := "abc.com"
	url := AddPrefixHttp(inputUrl)
	if strings.Compare(inputUrl, url) == 0 && strings.Compare(url, "http://abc.com") != 0 {
		t.Log("Prefix addition failed")
		t.Fail()
	}

	inputUrl = "http://google.com"
	url = AddPrefixHttp(inputUrl)
	if strings.Compare(inputUrl, url) != 0 {
		t.Log("Prefix addition failed")
		t.Fail()
	}
}
