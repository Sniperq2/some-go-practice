package main

import (
	"flag"
	"log"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func isFlag(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func flagChecker() {
	if !isFlag("from") {
		log.Fatal("Please provide a path to file to copy from")
	}

	if !isFlag("to") {
		log.Fatal("Please provide a filename to copy into")
	}
}

func main() {
	flag.Parse()
	flagChecker()

	if err := Copy(from, to, offset, limit); err != nil {
		log.Fatal(err)
	}
}
