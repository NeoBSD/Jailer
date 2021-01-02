package jailer

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Service ...
type Service struct {
	Label       string            `json:"label,omitempty" yaml:"label,omitempty"`
	Image       string            `json:"image,omitempty" yaml:"image,omitempty"`
	Ports       []string          `json:"ports,omitempty" yaml:"ports,omitempty"`
	Volumes     []string          `json:"volumes,omitempty" yaml:"volumes,omitempty"`
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Command     string            `json:"command,omitempty" yaml:"command,omitempty"`
	CPUPercent  int               `json:"cpu_percent,omitempty" yaml:"cpu_percent,omitempty"`
	CPUThreads  int               `json:"cpu_threads,omitempty" yaml:"cpu_threads,omitempty"`
}

// Compose ...
type Compose struct {
	Version  string    `json:"version" yaml:"version"`
	Services []Service `json:"services" yaml:"services"`
}

// ReadComposeFile reads and parses a `jailer-compose.yml` file.
func ReadComposeFile(path string) (*Compose, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	compose := &Compose{}
	if err := yaml.Unmarshal(buffer, &compose); err != nil {
		return nil, err
	}

	return compose, nil
}

// Validate returns nil, if the validation was successful.
func (compose *Compose) Validate() error {
	isInSlice := func(list []string, value string) bool {
		for _, v := range list {
			if v == value {
				return true
			}
		}
		return false
	}

	// Version
	if len(compose.Version) == 0 {
		return fmt.Errorf("no version specified")

	}
	// Unique labels
	knownLabels := []string{}
	for _, service := range compose.Services {
		label := service.Label
		if isInSlice(knownLabels, label) {
			return fmt.Errorf("label used twice: '%s'", label)
		}
		knownLabels = append(knownLabels, label)
	}
	return nil
}
