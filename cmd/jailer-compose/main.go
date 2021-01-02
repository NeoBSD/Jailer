package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

// JailerCompose ...
type JailerCompose struct {
	Version  string    `json:"version" yaml:"version"`
	Services []Service `json:"services" yaml:"services"`
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	path := "cmd/jailer-compose/testdata/jailer-compose.yml"
	compose, err := parseComposeFile(path)
	if err != nil {
		return err
	}

	// fmt.Println(compose)

	js, err := json.Marshal(compose)
	fmt.Println(string(js))
	if err != nil {
		return err
	}

	return nil
}

func parseComposeFile(path string) (JailerCompose, error) {
	buffer, err := ioutil.ReadFile(path)
	if err != nil {
		return JailerCompose{}, err
	}

	compose := JailerCompose{}
	if err := yaml.Unmarshal(buffer, &compose); err != nil {
		return JailerCompose{}, err
	}

	return compose, nil
}
