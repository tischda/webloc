package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	path := "Spring4TW! | Josh Long | Talk Video | Parleys.com.webloc"
	expected := "Spring4TW! - Josh Long - Talk Video - Parleys.com.url"
	newPath := convertPath(path)
	if newPath != expected {
		t.Errorf("Expected: %q, actual: %q", expected, newPath)
	}
}

func TestMainPokenv(t *testing.T) {
	args := []string{"-version"}
	os.Args = append(os.Args, args...)

	expected := fmt.Sprintf("webloc version %s\n", version)
	actual := captureOutput(main)

	if expected != actual {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

// captures Stdout and returns output of function f()
func captureOutput(f func()) string {
	// redirect output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	// reset output again
	w.Close()
	os.Stdout = old

	captured, _ := ioutil.ReadAll(r)
	return string(captured)
}
