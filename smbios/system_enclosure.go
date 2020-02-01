package smbios

import "github.com/digitalocean/go-smbios/smbios"

type SystemEnclosureStructure struct {
	smbios.Structure
}

func (s Smbios) SystemEnclosure() SystemEnclosureStructure {
	return s.SystemEnclosureStructure
}

func (s SystemEnclosureStructure) Manufacturer() string {
	return get(s.Structure, 0)
}

func (s SystemEnclosureStructure) Version() string {
	return get(s.Structure, 1)
}

func (s SystemEnclosureStructure) SerialNumber() string {
	return get(s.Structure, 2)
}

func (s SystemEnclosureStructure) AssetTagNumber() string {
	return get(s.Structure, 3)
}

func (s SystemEnclosureStructure) SKUNumber() string {
	return get(s.Structure, 4)
}
