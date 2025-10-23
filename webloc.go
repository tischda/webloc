package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"howett.net/plist"
)

// process handles the conversion of a single webloc file to a Windows .url file.
// It decodes the URL from the webloc file, prints it, and optionally creates the
// new .url file and deletes the original based on the configuration.
func process(path string, cfg *Config) {
	url := decode(path)

	if cfg.verbose {
		newPath := convertPath(path)
		fmt.Printf("\npath: %s --> %s\nurl: ", path, newPath)
	}
	if !cfg.quiet {
		fmt.Println(url)
	}
	if !cfg.noop {
		newPath := convertPath(path)
		writeUrl(newPath, url)
		if cfg.delete {
			err := os.Remove(path)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// decode extracts the URL from a macOS webloc file by parsing its plist format.
// It returns the URL string contained within the webloc file.
func decode(path string) string {
	var data weblocHeader

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // nolint:errcheck

	decoder := plist.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data.URL
}

// convertPath transforms a webloc file path to a corresponding .url file path.
// It changes the file extension from .webloc to .url and replaces any forbidden
// characters in Windows filenames with underscores.
func convertPath(path string) string {
	// replace filename extension
	newPath := path[:len(path)-len(".webloc")] + ".url"

	// replace forbidden characters
	r := strings.NewReplacer("|", "_", ":", "_", "?", "_", "<", "_", ">", "_", "*", "_", "\"", "_")
	newPath = r.Replace(newPath)

	return newPath
}

// writeUrl creates a Windows .url file at the specified path containing the given URL.
func writeUrl(path string, url string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // nolint:errcheck

	_, err = f.WriteString("[InternetShortcut]\nURL=" + url)
	if err != nil {
		log.Fatal(err)
	}
}
