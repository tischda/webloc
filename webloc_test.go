package main

import (
	"testing"
)

func TestDecodeBinary(t *testing.T) {
	expected := "http://www.kekaosx.com/en/"
	url := decode("data/binary-plist.webloc")
	if url != expected {
		t.Errorf("Expected: %q, actual: %q", expected, url)
	}
}

func TestDecodeXML(t *testing.T) {
	expected := "http://coffeescript.org/"
	url := decode("data/xml-content.webloc")
	if url != expected {
		t.Errorf("Expected: %q, actual: %q", expected, url)
	}
}

func TestConvertPath(t *testing.T) {
	path := "examples\\Spring4TW! | Josh Long | Talk Video : Parleys? <placeholder> and \"quote\".com.webloc"
	expected := "examples\\Spring4TW! _ Josh Long _ Talk Video _ Parleys_ _placeholder_ and _quote_.com.url"
	newPath := convertPath(path)
	if newPath != expected {
		t.Errorf("Expected: %q, actual: %q", expected, newPath)
	}
}
