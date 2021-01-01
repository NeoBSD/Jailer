package freebsd

import "strings"

// GetHostname returns general system infos
func GetHostname() (string, error) {
	out, err := getCommandOutput("hostname")
	if err != nil {
		return "", err
	}

	return strings.Trim(out, " \t\r\n"), nil
}
