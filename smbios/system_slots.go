package smbios

import "github.com/digitalocean/go-smbios/smbios"

type SystemSlotsStructure struct {
	smbios.Structure
}

func (s Smbios) SystemSlots() SystemSlotsStructure {
	return s.SystemSlotsStructure
}

func (s SystemSlotsStructure) SlotDesignation() string {
	return get(s.Structure, 0)
}
