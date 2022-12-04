package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdRun := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	cmdRun.Stderr = os.Stderr
	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout

	for key, value := range env {
		if value.NeedRemove {
			os.Unsetenv(key)
			continue
		}

		os.Setenv(key, value.Value)
	}

	cmdRun.Env = os.Environ()

	if err := cmdRun.Start(); err != nil {
		return -1
	}

	var comandErrr *exec.ExitError
	retCode := 0
	if err := cmdRun.Wait(); err != nil {
		if errors.As(err, &comandErrr) {
			retCode = comandErrr.ExitCode()
		}
	}
	return retCode
}
