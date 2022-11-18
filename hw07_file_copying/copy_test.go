package main

import (
	"testing"
)

func TestCopy(_ *testing.T) {
	Copy("./testdata/input.txt", "out.txt", 0, 100)
}
