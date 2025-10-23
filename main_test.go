//go:build windows

package main

import (
	"flag"
	"os"
	"testing"
)

func TestInitFlags(t *testing.T) {
	// Save original command line and reset flags
	originalArgs := os.Args
	originalCommandLine := flag.CommandLine

	defer func() {
		os.Args = originalArgs
		flag.CommandLine = originalCommandLine
	}()

	// Create a new flag set for this test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Test initFlags() function
	cfg := initFlags()

	// Test default values
	if cfg.help != false {
		t.Errorf("Expected help default to be false, got %v", cfg.help)
	}
	if cfg.version != false {
		t.Errorf("Expected version default to be false, got %v", cfg.version)
	}

	// Test that flags can be parsed
	testArgs := []string{
		"progname",
		"-?",
		"-v",
	}

	// Reset flag set and reinitialize
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cfg = initFlags()

	// Parse test arguments
	err := flag.CommandLine.Parse(testArgs[1:])
	if err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	// Verify flags were set correctly
	if !cfg.version {
		t.Error("Expected version flag to be true")
	}
}
