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

