package jail

// Jail represents a single jail config
type Jail struct {
	Name       string `json:"name"`
	Hostname   string `json:"hostname"`
	IP         string `json:"ip"`
	Path       string `json:"path"`
	MountDevFS bool   `json:"mount_devfs"`
	ExecStart  string `json:"exec_start"`
	ExecStop   string `json:"exec_stop"`
}
