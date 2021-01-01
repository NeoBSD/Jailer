package filesystem

import (
	"os/exec"
)

// NewDataset creates a zfs dataset for jail storage
func NewDataset(path string) error {
	return exec.Command("zfs", "create", path).Wait()
}
