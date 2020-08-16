package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents a set of configuration variables.
type Config struct {
	// The version of the configuration manifest.
	Version string `json:"version"`

	// List of devices and their information.
	Devices []Device `json:"devices"`
}

// Device represent a single Hikvision device.
type Device struct {
	// Friendly-name of the device.
	Name string `json:"name"`

	// Host address of the device.
	Host string `json:"host"`

	// Username for accessing the device.
	Username string `json:"username"`

	// Password for accessing the device.
	Password string `json:"password"`
}

// Parse parses yaml-encoded data in bytes.
func Parse(data []byte) (*Config, error) {
	var conf *Config
	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// ParseFile parses yaml-encoded data from a file.
func ParseFile(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Parse(b)
}
