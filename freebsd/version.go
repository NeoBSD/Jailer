package freebsd

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

// Version details a FreeBSD version
type Version struct {
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Patch   int    `json:"patch"`
	Comment string `json:"comment"`
}

// GetVersionInfo returns the freebsd-version command output
func GetVersionInfo() (Version, error) {
	var buffer bytes.Buffer
	cmd := exec.Command("freebsd-version")
	cmd.Stdout = &buffer
	cmd.Stderr = &buffer

	// Output will be something like: 12.2-RELEASE-p1
	if err := cmd.Run(); err != nil {
		return Version{}, err
	}

	version := Version{}

	versionString := string(buffer.Bytes())
	majorSplit := strings.Split(versionString, ".")
	majorString := majorSplit[0]
	major, err := strconv.Atoi(majorString)
	if err != nil {
		return Version{}, err
	}
	version.Major = major

	minorSplits := strings.Split(majorSplit[1], "-")
	minor, err := strconv.Atoi(minorSplits[0])
	if err != nil {
		return Version{}, err
	}
	version.Minor = minor

	if len(minorSplits) > 2 {
		str := minorSplits[2][1:]
		str = strings.Trim(str, " \n\r\t")
		patch, err := strconv.Atoi(str)
		if err != nil {
			return Version{}, err
		}
		version.Patch = patch
	}

	version.Comment = strings.Trim(minorSplits[1], " \t\n\r")

	return version, nil
}
