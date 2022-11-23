package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFine(t *testing.T) {
	err := Copy("./testdata/input.txt", "out.txt", 0, 100)
	assert.Nil(t, err, "must copying fine")
}

func TestCopyFileSizeIsGreater(t *testing.T) {
	if err := Copy("./testdata/input.txt", "out.txt", 7000, 100); err != nil {
		assert.ErrorContains(t, err, "offset exceeds file size", "file size is greater then offset")
	}
}

func TestCopyForUnsupportedFiles(t *testing.T) {
	if err := Copy("C:\\msys64\\dev\\stderr", "out.txt", 0, 100); err != nil {
		assert.ErrorContains(t, err, "unsupported file", "unsupported file test")
	}
}
