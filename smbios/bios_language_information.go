package smbios

import "github.com/digitalocean/go-smbios/smbios"

type BIOSLanguageInformationStructure struct {
	smbios.Structure
}

func (s Smbios) BIOSLanguageInformation() BIOSLanguageInformationStructure {
	return s.BIOSLanguageInformationStructure
}

func (s BIOSLanguageInformationStructure) CurrentLanguage() string {
	return get(s.Structure, 0)
}
