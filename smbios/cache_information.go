package smbios

import "github.com/digitalocean/go-smbios/smbios"

type CacheInformationStructure struct {
	smbios.Structure
}

func (s Smbios) CacheInformation() CacheInformationStructure {
	return s.CacheInformationStructure
}

func (s CacheInformationStructure) SocketDesignation() string {
	return get(s.Structure, 0)
}
