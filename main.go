package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/tobiashienzsch/jailer/cmd"
	"github.com/tobiashienzsch/jailer/runtime"
)

// These variables get set during link time. See Makefile
var (
	hostOS string
	commit string
	date   string
)

func init() {

	// Set build info
	runtime.BuildCommit = commit
	runtime.BuildDate = date
	runtime.BuildOS = hostOS

	// Log as JSON instead of the default ASCII formatter.
	// logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only logrus the warning severity or above.
	logrus.SetLevel(logrus.WarnLevel)

}

func main() {
	cmd.Execute()
}
