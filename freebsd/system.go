package freebsd

// System details a FreeBSD version
type System struct {
	Hostname          string             `json:"hostname"`
	Machine           string             `json:"machine"`
	MachineArch       string             `json:"machine_arch"`
	Hardware          Hardware           `json:"hardware"`
	Version           Version            `json:"version"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces"`
	PFCTL             PFCTL              `json:"pfctl"`
}

// GetSystemInfo returns general system infos
func GetSystemInfo() (*System, error) {
	hostname, err := GetHostname()
	if err != nil {
		return nil, err
	}

	machine, err := getCommandOutput("uname", "-m")
	if err != nil {
		return nil, err
	}

	machineArch, err := getCommandOutput("uname", "-p")
	if err != nil {
		return nil, err
	}

	hardware, err := GetHardwareInfo()
	if err != nil {
		return nil, err
	}

	version, err := GetVersionInfo()
	if err != nil {
		return nil, err
	}

	system := System{
		Hostname:    hostname,
		Machine:     machine,
		MachineArch: machineArch,
		Hardware:    hardware,
		Version:     version,
	}

	return &system, nil
}
