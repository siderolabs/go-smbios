package smbios

import "github.com/digitalocean/go-smbios/smbios"

type GroupAssociationsStructure struct {
	smbios.Structure
}

func (s Smbios) GroupAssociations() GroupAssociationsStructure {
	return s.GroupAssociationsStructure
}

func (s GroupAssociationsStructure) GroupName() string {
	return get(s.Structure, 0)
}
