package smbios

import "github.com/digitalocean/go-smbios/smbios"

type PortConnectorInformationStructure struct {
	smbios.Structure
}

func (s Smbios) PortConnectorInformation() PortConnectorInformationStructure {
	return s.PortConnectorInformationStructure
}

func (s PortConnectorInformationStructure) InternalReferenceDesignator() string {
	return get(s.Structure, 0)
}

func (s PortConnectorInformationStructure) ExternalReferenceDesignator() string {
	return get(s.Structure, 2)
}
