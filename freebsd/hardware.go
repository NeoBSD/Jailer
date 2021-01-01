package freebsd

import (
	"os/exec"
	"strconv"
	"strings"
)

// Hardware ...
type Hardware struct {
	Model  string `json:"model"`
	NumCPU int    `json:"ncpu"`
}

// GetHardwareInfo returns `sysctl -a | grep hw.model`
func GetHardwareInfo() (Hardware, error) {
	hardware := Hardware{}

	// CPU model
	out, err := GetPipeCommandOutput(
		exec.Command("sysctl", "-a"),
		exec.Command("grep", "hw.model:"),
	)
	if err != nil || len(out) == 0 {
		return Hardware{}, err
	}
	out = strings.TrimPrefix(out, "hw.model: ")
	out = strings.Trim(out, " \t\r\n")
	hardware.Model = out

	// Num CPU
	out, err = GetPipeCommandOutput(
		exec.Command("sysctl", "-a"),
		exec.Command("grep", "hw.ncpu:"),
	)
	if err != nil || len(out) == 0 {
		return Hardware{}, err
	}
	out = strings.TrimPrefix(out, "hw.ncpu: ")
	out = strings.Trim(out, " \t\r\n")
	ncpu, err := strconv.Atoi(out)
	if err != nil {
		return Hardware{}, err
	}
	hardware.NumCPU = ncpu

	return hardware, nil
}
