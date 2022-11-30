package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if len(dir) == 0 {
		return nil, fmt.Errorf("please set directory")
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("no directory found")
	}

	infos := make([]string, 0, len(entries))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			log.Fatalf("errror rading files")
		}
		infos = append(infos, info.Name())
	}

	envList := Environment{}
	for _, i := range infos {
		file, err := os.Open(path.Join(dir, i))
		if err != nil {
			return nil, fmt.Errorf("")
		}
		reader := bufio.NewReader(file)

		value, _, err := reader.ReadLine()
		if err != nil {
			return nil, fmt.Errorf("")
		}

		outValue := string(value)
		outValue = strings.Trim(outValue, " \t")
		resultValue := string(bytes.ReplaceAll([]byte(outValue), []byte{0x00}, []byte("\n")))

		if len(resultValue) == 0 {
			continue
		}

		newValue := EnvValue{
			Value:      resultValue,
			NeedRemove: false,
		}
		envList[i] = newValue
		fmt.Printf("%#v", envList)
	}
	return envList, nil
}
