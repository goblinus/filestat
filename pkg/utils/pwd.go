package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func Pwd() (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("pwd")
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
