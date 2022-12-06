package main

import (
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	timeout string
)

func main() {
	flag.StringVar(&timeout, "timeout", "10s", "timeout")
	flag.Parse()
	host := os.Args[2]
	port := os.Args[3]
	if len(host) == 0 {
		log.Fatalf("Please set host")
	}
	if len(port) == 0 {
		log.Fatalf("Please set port")
	}

	fmt.Printf("%s %s %s", timeout, host, port)
}
