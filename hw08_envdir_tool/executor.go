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
		if len(value.Value) == 0 {
			value.NeedRemove = true
			os.Unsetenv(key)
			continue
		}

		os.Setenv(key, value.Value)
	}

	cmdRun.Env = os.Environ()
	var commandErr *exec.ExitError
	retCode := 0 // если запуск пройдёт гладко, то так и останется 0
	if err := cmdRun.Run(); err != nil {
		// Из документации:
		// If the command starts but does not complete successfully, the error is of type *ExitError.
		if errors.As(err, &commandErr) {
			// нам сама ошибка не нуждан - только ExitCode
			retCode = commandErr.ExitCode()
		}
	}
	return retCode
}
