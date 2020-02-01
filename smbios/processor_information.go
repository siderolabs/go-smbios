package smbios

import "github.com/digitalocean/go-smbios/smbios"

type ProcessorInformationStructure struct {
	smbios.Structure
}

func (s Smbios) ProcessorInformation() ProcessorInformationStructure {
	return s.ProcessorInformationStructure
}

func (s ProcessorInformationStructure) SocketDesignation() string {
	return get(s.Structure, 0)
}

func (s ProcessorInformationStructure) ProcessorManufacturer() string {
	return get(s.Structure, 1)
}

func (s ProcessorInformationStructure) ProcessorVersion() string {
	return get(s.Structure, 2)
}

func (s ProcessorInformationStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

func (s ProcessorInformationStructure) AssetTag() string {
	return get(s.Structure, 4)
}

func (s ProcessorInformationStructure) PartNumber() string {
	return get(s.Structure, 5)
}
