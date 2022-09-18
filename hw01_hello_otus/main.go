package main

import (
	"fmt"
	"os"

	"golang.org/x/example/stringutil"
)

func main() {
	reversed := stringutil.Reverse("Hello, OTUS!")
	fmt.Fprintln(os.Stdout, reversed)
}
