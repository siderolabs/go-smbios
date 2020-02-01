package smbios

import "github.com/digitalocean/go-smbios/smbios"

type BIOSInformationStructure struct {
	smbios.Structure
}

func (s Smbios) BIOSInformation() BIOSInformationStructure {
	return s.BIOSInformationStructure
}

func (s BIOSInformationStructure) Vendor() string {
	return get(s.Structure, 0)
}

func (s BIOSInformationStructure) Version() string {
	return get(s.Structure, 1)
}

func (s BIOSInformationStructure) ReleaseDate() string {
	return get(s.Structure, 2)
}
