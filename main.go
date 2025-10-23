package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// https://goreleaser.com/cookbooks/using-main.version/
var (
	name    string
	version string
	date    string
	commit  string
)

// flags
type Config struct {
	delete  bool
	noop    bool
	quiet   bool
	verbose bool
	help    bool
	version bool
}

func initFlags() *Config {
	cfg := &Config{}
	flag.BoolVar(&cfg.delete, "d", true, "")
	flag.BoolVar(&cfg.delete, "delete", true, "delete .webloc files after conversion")
	flag.BoolVar(&cfg.noop, "n", false, "")
	flag.BoolVar(&cfg.noop, "noop", false, "decode urls, but do not change file system")
	flag.BoolVar(&cfg.quiet, "q", false, "")
	flag.BoolVar(&cfg.quiet, "quiet", false, "suppress non-error output")
	flag.BoolVar(&cfg.verbose, "vv", false, "")
	flag.BoolVar(&cfg.verbose, "verbose", false, "verbose logging")
	flag.BoolVar(&cfg.help, "?", false, "")
	flag.BoolVar(&cfg.help, "help", false, "displays this help message")
	flag.BoolVar(&cfg.version, "v", false, "")
	flag.BoolVar(&cfg.version, "version", false, "print version and exit")
	return cfg
}

type weblocHeader struct {
	URL string `plist:"URL"`
}

func main() {
	log.SetFlags(0)
	cfg := initFlags()
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: "+name+` [OPTIONS] <path> 

Recursively converts all .webloc files in path to .url files.

OPTIONS:
  -d, --delete
		delete .webloc files after conversion (default: true)
  -n, --noop
		decode urls, but do not change files
  -q, --quiet
		suppress non-error output
  -vv, --verbose
		enable verbose logging
  -?, --help
        display this help message
  -v, --version
        print version and exit

EXAMPLES:`)

		fmt.Fprintln(os.Stderr, "\n  $ "+name+` --noop data

    http://www.kekaosx.com/en/
    https://www.maketecheasier.com/fix-home-end-button-for-external-keyboard-mac/
    http://coffeescript.org/`)
	}
	flag.Parse()

	if flag.Arg(0) == "version" || cfg.version {
		fmt.Printf("%s %s, built on %s (commit: %s)\n", name, version, date, commit)
		return
	}

	if cfg.help {
		flag.Usage()
		return
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// process path recursively
	walkpath := func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasPrefix(f.Name(), ".") {
			matched, err := filepath.Match("*.webloc", f.Name())
			if err != nil {
				return err
			}
			if matched {
				process(path, cfg)
			}
		}
		return nil
	}

	err := filepath.Walk(flag.Arg(0), walkpath)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.noop && !cfg.quiet {
		fmt.Println("\n--noop: No file changes were made.")
	}
}
