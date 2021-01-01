package jailer

import (
	"os/exec"
)

// ZFSDataset ...
type ZFSDataset struct {
	Path string `json:"path"`
}

// NewZFSDataset creates a zfs dataset for jail storage
func NewZFSDataset(path string) (*ZFSDataset, error) {
	err := exec.Command("zfs", "create", path).Wait()
	return &ZFSDataset{Path: path}, err
}
