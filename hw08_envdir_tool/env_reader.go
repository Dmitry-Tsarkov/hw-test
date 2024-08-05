package main

import (
	"os"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func cleanString(s string) string {
	var cleanedLine strings.Builder
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch != '\x00' {
			cleanedLine.WriteRune(rune(ch))
		} else {
			cleanedLine.WriteRune('\n')
		}
	}
	return cleanedLine.String()
}

func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileData, err := os.ReadFile(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		lines := strings.Split(string(fileData), "\n")
		value := cleanString(lines[0])

		env[file.Name()] = EnvValue{
			Value:      value,
			NeedRemove: len(value) == 0,
		}
	}

	return env, nil
}
