package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmd []string, env Environment) (string, int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	for key, value := range env {
		if value.NeedRemove {
			os.Unsetenv(key)
		} else {
			// cleanValue := strings.TrimSpace(value.Value)
			// os.Setenv(key, cleanValue)
			cleanValue := strings.Replace(strings.TrimRight(value.Value, " \n"), "\n", "\n", -1)
			os.Setenv(key, cleanValue)
		}
	}

	out, err := command.CombinedOutput()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return string(out), exitError.ExitCode()
		}
		return "", 1
	}

	return string(out), 0
}
