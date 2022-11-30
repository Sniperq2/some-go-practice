package main

import (
	"log"
	"os"
)

func main() {
	e, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatalf("%v", err)
	}

	code := RunCmd(os.Args[2:], e)
	os.Exit(code)
}
