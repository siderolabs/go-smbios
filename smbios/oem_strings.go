package smbios

import "github.com/digitalocean/go-smbios/smbios"

type OEMStringsStructure struct {
	smbios.Structure
}

func (s Smbios) OEMStrings() OEMStringsStructure {
	return s.OEMStringsStructure
}
