package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdRun := exec.Command(cmd[0], cmd[1:]...)

	cmdRun.Stderr = os.Stderr
	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout

	additionalEnv := make([]string, 0)
	for key, value := range env {
		additionalEnv = append(additionalEnv, fmt.Sprintf("%s=%s", key, value.Value))
	}

	cmdRun.Env = append(os.Environ(), additionalEnv...)

	if err := cmdRun.Run(); err != nil {
		return 1
	}

	return cmdRun.ProcessState.ExitCode()
}
