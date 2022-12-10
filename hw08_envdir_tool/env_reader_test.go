package main

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("part env variables", func(t *testing.T) {
		tmpDir, err := os.MkdirTemp("testdata", "hw_08_test")
		if err != nil {
			log.Fatalf("could not create temp directory: %s", err)
		}
		defer os.RemoveAll(tmpDir)

		barFile, err := os.Create(filepath.Join(tmpDir, "BAR"))
		if err != nil {
			log.Fatal(err)
		}
		barFile.WriteString("bar\nPLEASE IGNORE SECOND LINE\n")

		helloFile, err := os.Create(filepath.Join(tmpDir, "HELLO"))
		if err != nil {
			log.Fatal(err)
		}
		helloFile.WriteString("\"hello\"")

		cleanup := func() {
			barFile.Close()
			helloFile.Close()
			os.Remove("BAR")
			os.Remove("HELLO")
		}

		defer cleanup()

		envMap := Environment{}
		envMap["BAR"] = EnvValue{Value: "bar", NeedRemove: false}
		envMap["HELLO"] = EnvValue{Value: "\"hello\"", NeedRemove: false}
		funcEnv, err := ReadDir(tmpDir)

		require.NoError(t, err)
		require.Equal(t, envMap, funcEnv, "environments are not the same")
	})
}
