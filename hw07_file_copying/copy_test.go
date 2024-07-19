package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// Test case struct
type testCase struct {
	name         string
	fromFile     string
	toFile       string
	offset       int64
	limit        int64
	expectedFile string
	expectedErr  error
}

// Define test cases
var tests = []testCase{
	{
		name:         "copy entire file",
		fromFile:     "input.txt",
		toFile:       "out_offset0_limit0.txt",
		offset:       0,
		limit:        0,
		expectedFile: "out_offset0_limit0.txt",
		expectedErr:  nil,
	},
	{
		name:         "copy with limit",
		fromFile:     "input.txt",
		toFile:       "out_offset0_limit10.txt",
		offset:       0,
		limit:        10,
		expectedFile: "out_offset0_limit10.txt",
		expectedErr:  nil,
	},
	{
		name:         "copy with offset",
		fromFile:     "input.txt",
		toFile:       "out_offset100_limit1000.txt",
		offset:       100,
		limit:        1000,
		expectedFile: "out_offset100_limit1000.txt",
		expectedErr:  nil,
	},
	{
		name:         "file does not exist",
		fromFile:     "nonexistent.txt",
		toFile:       "out_nonexistent.txt",
		offset:       0,
		limit:        0,
		expectedFile: "",
		expectedErr:  os.ErrNotExist,
	},
}

func TestCopy(t *testing.T) {
	// Define the test data directory
	testDataDir := "testData"

	// Define a helper function to read file contents
	readFileContents := func(path string) ([]byte, error) {
		return os.ReadFile(path)
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file for the destination
			tempFile, err := os.CreateTemp("", "copy_test_*")
			if err != nil {
				t.Fatalf("unable to create temp file: %v", err)
			}
			defer os.Remove(tempFile.Name())
			tempFile.Close()

			err = Copy(filepath.Join(testDataDir, tt.fromFile), tempFile.Name(), tt.offset, tt.limit)

			if err != nil {
				if tt.expectedErr == nil || !os.IsNotExist(err) {
					t.Errorf("unexpected error: %v", err)
				}
			} else {
				expectedContents, err := readFileContents(filepath.Join(testDataDir, tt.expectedFile))
				if err != nil {
					t.Fatalf("unable to read expected file: %v", err)
				}

				actualContents, err := readFileContents(tempFile.Name())
				if err != nil {
					t.Fatalf("unable to read actual file: %v", err)
				}

				if !bytes.Equal(expectedContents, actualContents) {
					t.Errorf("contents do not match: expected %s, got %s", expectedContents, actualContents)
				}
			}
		})
	}
}
