package main

import (
	"flag"
	"fmt"
	"howett.net/plist"
	"os"
	"path/filepath"
	"strings"
)

const version string = "1.0.0"

type weblocHeader struct {
	URL string `plist:"URL"`
}

var delete bool
var noop bool

func init() {
	flag.BoolVar(&delete, "delete", false, "delete .webloc files after conversion")
	flag.BoolVar(&noop, "noop", false, "decode urls, but do not change file system")
}

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] path\n  path: the path to process\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *showVersion {
		fmt.Println("webloc version", version)
		return
	}

	filepath.Walk(flag.Arg(0), walkpath)
}

func walkpath(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	matched, err := filepath.Match("*.webloc", f.Name())
	if err != nil {
		return err
	}
	if matched {
		process(path)
	}
	return nil
}

func process(path string) {
	url := decode(path)
	fmt.Println(url)

	if !noop {
		newPath := convertPath(path)
		writeUrl(newPath, url)
		if delete {
			err := os.Remove(path)
			check(err)
		}
	}
}

func decode(path string) string {
	var data weblocHeader

	f, err := os.Open(path)
	check(err)
	defer f.Close()

	decoder := plist.NewDecoder(f)
	check(decoder.Decode(&data))
	return data.URL
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertPath(path string) string {
	newPath := path[:len(path)-len(".webloc")] + ".url"
	newPath = strings.Replace(newPath, "|", "-", -1)
	return newPath
}

func writeUrl(path string, url string) {
	f, err := os.Create(path)
	check(err)
	defer f.Close()

	f.WriteString("[InternetShortcut]\nURL=" + url)
}
