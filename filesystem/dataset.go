package filesystem

import (
	"fmt"
	"os/exec"
)

// NewDataset creates a zfs dataset for container storage
func NewDataset(path string) error {
	// Setup external zfs list
	zfsCMD := "zfs"
	c := exec.Command(zfsCMD, "create", path)

	// Exec external zfs list
	stdout, err := c.Output()
	if err != nil {
		return fmt.Errorf("Error while executing external command: %v", err.Error())
	}

	fmt.Print(string(stdout))
	return nil
}
