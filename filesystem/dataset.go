package filesystem

import (
	"github.com/sirupsen/logrus"

	"fmt"
	"os/exec"
)

// NewDataset creates a zfs dataset for container storage
func NewDataset(path string) error {
	// Setup external zfs list
	zfsCMD := "zfs"
	c := exec.Command(zfsCMD, "create", path)

	logrus.Warn(c.Args)
	if err := c.Start(); err != nil {
		logrus.Fatal(err)
	}

	// Wait for the process to finish
	done := make(chan error, 1)
	go func() {
		done <- c.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("process finished with error = %v", err)
		}
		fmt.Println("process finished successfully")
	}

	return nil
}
