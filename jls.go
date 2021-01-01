package jailer

import (
	"os/exec"
	"strconv"
	"strings"
)

// JLS manages the `jls` command on FreeBSD
type JLS struct {
	Path string `json:"path"`
}

// GetActiveJails returns a slice of all running jails
func (jls *JLS) GetActiveJails() ([]Jail, error) {
	output, err := exec.Command(jls.Path, "-n").Output()
	if err != nil {
		return nil, err
	}

	jails, err := parseJLS(string(output))
	if err != nil {
		return nil, err
	}

	return jails, nil
}

func parseJLS(str string) ([]Jail, error) {
	jails := []Jail{}
	lines := strings.Split(str, "\n")
	for _, line := range lines {
		// Split at spaces
		items := strings.Split(line, " ")

		// Ignore empty lines
		if len(strings.Trim(line, " \n\r\t")) == 0 {
			continue
		}

		j := Jail{}

		// Mapping
		intValues := []struct {
			ID    string
			Value *int
		}{
			{ID: "jid", Value: &j.JID},
			{ID: "parent", Value: &j.Parent},
			{ID: "host.hostid", Value: &j.HostID},
			{ID: "osreldate", Value: &j.OSReleaseDate},
			{ID: "devfs_ruleset", Value: &j.DevFSRuleset},
			{ID: "enforce_statfs", Value: &j.EnforceStatFS},
			{ID: "securelevel", Value: &j.SecureLevel},
			{ID: "children.cur", Value: &j.ChildrenCurrent},
			{ID: "children.max", Value: &j.ChildrenMax},
			{ID: "cpuset.id", Value: &j.CPUSetID},
		}

		stringValues := []struct {
			ID    string
			Value *string
		}{
			{ID: "name", Value: &j.Name},
			{ID: "path", Value: &j.Path},
			{ID: "host.hostname", Value: &j.Hostname},
			{ID: "host.hostuuid", Value: &j.HostUUD},
			{ID: "osrelease", Value: &j.OSRelease},
		}

		boolValues := []struct {
			ID    string
			Value *bool
		}{
			{ID: "allow.nochflags", Value: &j.AllowNochFlags},
			{ID: "allow.nomlock", Value: &j.AllowNoMLock},
			{ID: "allow.nomount", Value: &j.AllowNoMount},
			{ID: "allow.mount.nodevfs", Value: &j.AllowMountNoDevFS},
			{ID: "allow.mount.noprocfs", Value: &j.AllowMountNoProcFS},
			{ID: "allow.mount.notmpfs", Value: &j.AllowMountNoTmpFS},
			{ID: "allow.mount.nozfs", Value: &j.AllowMountNoZFS},
			{ID: "allow.noquotas", Value: &j.AllowNoQuotas},
			{ID: "allow.noraw_sockets", Value: &j.AllowNoRawSockets},
			{ID: "allow.noread_msgbuf", Value: &j.AllowNoReadMsgBuf},
			{ID: "allow.reserved_ports", Value: &j.AllowReservedPorts},
			{ID: "allow.set_hostname", Value: &j.AllowSetHostname},
			{ID: "allow.nosocket_af", Value: &j.AllowNoSocketAF},
			{ID: "allow.nosysvipc", Value: &j.AllowNoSysVIPC},
		}

		// For each config item
		for _, item := range items {
			// Parse int values
			for _, v := range intValues {
				if strings.HasPrefix(item, v.ID) {
					item = strings.TrimPrefix(item, v.ID)
					item = strings.TrimPrefix(item, "=")
					num, err := strconv.Atoi(item)
					if err != nil {
						return nil, err
					}

					*v.Value = num
				}
			}

			// Parse string values
			for _, v := range stringValues {
				if strings.HasPrefix(item, v.ID) {
					item = strings.TrimPrefix(item, v.ID)
					item = strings.TrimPrefix(item, "=")
					*v.Value = strings.Trim(item, " \n\t\r")
				}
			}

			// Parse boolean values
			for _, v := range boolValues {
				if strings.HasPrefix(item, v.ID) {
					*v.Value = true
				}
			}

		}

		jails = append(jails, j)
	}

	return jails, nil
}
