package smbios

import (
	"fmt"
	"strings"
	"sync"

	"github.com/digitalocean/go-smbios/smbios"
)

type Smbios struct {
	Version    string
	Structures []*smbios.Structure

	BIOSInformationStructure            BIOSInformationStructure
	SystemInformationStructure          SystemInformationStructure
	BaseboardInformationStructure       BaseboardInformationStructure
	SystemEnclosureStructure            SystemEnclosureStructure
	ProcessorInformationStructure       ProcessorInformationStructure
	CacheInformationStructure           CacheInformationStructure
	PortConnectorInformationStructure   PortConnectorInformationStructure
	SystemSlotsStructure                SystemSlotsStructure
	OEMStringsStructure                 OEMStringsStructure
	SystemConfigurationOptionsStructure SystemConfigurationOptionsStructure
	BIOSLanguageInformationStructure    BIOSLanguageInformationStructure
	GroupAssociationsStructure          GroupAssociationsStructure
}

func New() (*Smbios, error) {
	rc, ep, err := smbios.Stream()
	if err != nil {
		return nil, fmt.Errorf("failed to open stream: %w", err)
	}

	defer rc.Close()

	s := &Smbios{}

	major, minor, rev := ep.Version()
	s.Version = fmt.Sprintf("%d.%d.%d", major, minor, rev)

	d := smbios.NewDecoder(rc)
	ss, err := d.Decode()
	if err != nil {
		return nil, fmt.Errorf("failed to decode structures: %w", err)
	}

	s.Structures = ss

	var wg sync.WaitGroup

	wg.Add(len(s.Structures))

	for _, structure := range s.Structures {
		go func(ss *smbios.Structure) {
			defer wg.Done()

			switch ss.Header.Type {
			case 0:
				s.BIOSInformationStructure = BIOSInformationStructure{Structure: *ss}
			case 1:
				s.SystemInformationStructure = SystemInformationStructure{Structure: *ss}
			case 2:
				s.BaseboardInformationStructure = BaseboardInformationStructure{Structure: *ss}
			case 3:
				s.SystemEnclosureStructure = SystemEnclosureStructure{Structure: *ss}
			case 4:
				s.ProcessorInformationStructure = ProcessorInformationStructure{Structure: *ss}
			case 5:
				// Obsolete.
			case 6:
				// Obsolete.
			case 7:
				s.CacheInformationStructure = CacheInformationStructure{Structure: *ss}
			case 8:
				s.PortConnectorInformationStructure = PortConnectorInformationStructure{Structure: *ss}
			case 9:
				s.SystemSlotsStructure = SystemSlotsStructure{Structure: *ss}
			case 10:
				// Obsolete.
			case 11:
				s.OEMStringsStructure = OEMStringsStructure{Structure: *ss}
			case 12:
				s.SystemConfigurationOptionsStructure = SystemConfigurationOptionsStructure{Structure: *ss}
			case 13:
				s.BIOSLanguageInformationStructure = BIOSLanguageInformationStructure{Structure: *ss}
			case 14:
				s.GroupAssociationsStructure = GroupAssociationsStructure{Structure: *ss}
			}
		}(structure)
	}

	wg.Wait()

	return s, nil
}

func get(s smbios.Structure, i int) string {
	unknown := "Unknown"

	if i >= len(s.Strings) {
		return unknown
	}

	// Convert to lowercase to address multiple formats:
	//   - "To Be Filled By O.E.M."
	//   - "To be filled by O.E.M."
	if strings.ToLower(s.Strings[i]) == "to be filled by o.e.m." {
		return unknown
	}

	return strings.TrimSpace(s.Strings[i])
}
