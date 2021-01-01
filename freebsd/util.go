package freebsd

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
)

func getCommandOutput(exe string, args ...string) (string, error) {
	var buffer bytes.Buffer
	cmd := exec.Command(exe, args...)
	cmd.Stdout = &buffer
	cmd.Stderr = &buffer
	if err := cmd.Run(); err != nil {
		return "", err
	}

	b := buffer.Bytes()
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)

	return strings.Trim(string(b), " \t\n\r"), nil
}

func GetPipeCommandOutput(c1, c2 *exec.Cmd) (string, error) {
	r, w := io.Pipe()
	c1.Stdout = w

	var buffer bytes.Buffer
	c2.Stdin = r
	c2.Stdout = &buffer

	if err := c1.Start(); err != nil {
		return "", nil
	}
	if err := c2.Start(); err != nil {
		return "", nil
	}
	if err := c1.Wait(); err != nil {
		return "", nil
	}
	if err := w.Close(); err != nil {
		return "", nil
	}
	if err := c2.Wait(); err != nil {
		return "", nil
	}

	return buffer.String(), nil
}
