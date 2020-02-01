package smbios

import "github.com/digitalocean/go-smbios/smbios"

type BaseboardInformationStructure struct {
	smbios.Structure
}

func (s Smbios) BaseboardInformation() BaseboardInformationStructure {
	return s.BaseboardInformationStructure
}

func (s BaseboardInformationStructure) Manufacturer() string {
	return get(s.Structure, 0)
}

func (s BaseboardInformationStructure) Product() string {
	return get(s.Structure, 1)
}

func (s BaseboardInformationStructure) Version() string {
	return get(s.Structure, 2)
}

func (s BaseboardInformationStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

func (s BaseboardInformationStructure) AssetTag() string {
	return get(s.Structure, 4)
}

func (s BaseboardInformationStructure) LocationInChassis() string {
	return get(s.Structure, 5)
}
