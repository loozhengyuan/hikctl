package config

import (
	"github.com/loozhengyuan/hikvision-sdk/hikvision"
)

// Apply applies the configuration to the device.
func (c *Config) Apply() error {
	for _, device := range c.Devices {
		_, err := hikvision.NewClient(device.Host, device.Username, device.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
