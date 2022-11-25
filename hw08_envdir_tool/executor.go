package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdRun := exec.Command(cmd[0], cmd[1:]...)

	cmdRun.Stderr = os.Stderr
	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout

	if err := cmdRun.Run(); err != nil {
		return -1
	}

	return cmdRun.ProcessState.ExitCode()
}
