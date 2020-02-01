package smbios

import "github.com/digitalocean/go-smbios/smbios"

type SystemConfigurationOptionsStructure struct {
	smbios.Structure
}

func (s Smbios) SystemConfigurationOptions() SystemConfigurationOptionsStructure {
	return s.SystemConfigurationOptionsStructure
}
