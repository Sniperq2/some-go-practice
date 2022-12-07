package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	ErrOtherFsError     = errors.New("other error, but not end of file")
	ErrSetDirectory     = errors.New("please set directory")
	ErrNoDirectoryFound = errors.New("no directory found")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func readParam(handle io.Reader) (string, error) {
	reader := bufio.NewReader(handle)
	value, _, err := reader.ReadLine()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return "", nil
		}
		return "", ErrOtherFsError
	}
	value = []byte(strings.TrimRight(string(value), " \t\x00"))
	value = bytes.ReplaceAll(value, []byte{0x00}, []byte("\n"))

	return string(value), nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if len(dir) == 0 {
		return nil, ErrSetDirectory
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrNoDirectoryFound
	}

	infos := make([]string, 0, len(entries))

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				return nil, os.ErrNotExist
			}
		}
		infos = append(infos, info.Name())
	}

	envList := Environment{}
	for _, i := range infos {
		if strings.Contains(i, "=") {
			continue
		}

		file, err := os.Open(path.Join(dir, i))
		if err != nil {
			return nil, fmt.Errorf("could not open file %s", i)
		}

		defer func() {
			if errClose := file.Close(); err != nil {
				err = errClose
			}
		}()

		resultValue, err := readParam(file)
		if err != nil {
			return nil, fmt.Errorf("could not read file with env variable %s", i)
		}

		newValue := EnvValue{
			Value:      resultValue,
			NeedRemove: false,
		}
		envList[i] = newValue
	}
	return envList, nil
}
