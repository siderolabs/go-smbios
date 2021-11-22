// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"fmt"
	"strings"

	"github.com/digitalocean/go-smbios/smbios"
)

// SMBIOS represents the Sysytem Management BIOS.
type SMBIOS struct { //nolint:govet
	Version struct {
		Major    int
		Minor    int
		Revision int
	}
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
	PhysicalMemoryArrayStructure        PhysicalMemoryArrayStructure
	MemoryDeviceStructure               MemoryDeviceStructure
}

// New initializes and returns a new `SMBIOS`
//
//nolint:gocyclo,cyclop
func New() (*SMBIOS, error) {
	rc, ep, err := smbios.Stream()
	if err != nil {
		return nil, fmt.Errorf("failed to open stream: %w", err)
	}

	//nolint: errcheck
	defer rc.Close()

	s := &SMBIOS{}

	s.Version.Major, s.Version.Minor, s.Version.Revision = ep.Version()

	d := smbios.NewDecoder(rc)

	ss, err := d.Decode()
	if err != nil {
		return nil, fmt.Errorf("failed to decode structures: %w", err)
	}

	s.Structures = ss

	for _, structure := range s.Structures {
		switch structure.Header.Type {
		case 0:
			s.BIOSInformationStructure = BIOSInformationStructure{Structure: structure}
		case 1:
			s.SystemInformationStructure = SystemInformationStructure{Structure: structure}
		case 2:
			s.BaseboardInformationStructure = BaseboardInformationStructure{Structure: structure}
		case 3:
			s.SystemEnclosureStructure = SystemEnclosureStructure{Structure: structure}
		case 4:
			s.ProcessorInformationStructure = ProcessorInformationStructure{Structure: structure}
		case 5:
			// Obsolete.
		case 6:
			// Obsolete.
		case 7:
			s.CacheInformationStructure = CacheInformationStructure{Structure: structure}
		case 8:
			s.PortConnectorInformationStructure = PortConnectorInformationStructure{Structure: structure}
		case 9:
			s.SystemSlotsStructure = SystemSlotsStructure{Structure: structure}
		case 10:
			// Obsolete.
		case 11:
			s.OEMStringsStructure = OEMStringsStructure{Structure: structure}
		case 12:
			s.SystemConfigurationOptionsStructure = SystemConfigurationOptionsStructure{Structure: structure}
		case 13:
			s.BIOSLanguageInformationStructure = BIOSLanguageInformationStructure{Structure: structure}
		case 14:
			s.GroupAssociationsStructure = GroupAssociationsStructure{Structure: structure}
		case 15:
			// Unimplemented.
		case 16:
			s.PhysicalMemoryArrayStructure = PhysicalMemoryArrayStructure{Structure: structure}
		case 17:
			s.MemoryDeviceStructure = MemoryDeviceStructure{Structure: structure}
		}
	}

	return s, nil
}

func get(s *smbios.Structure, i int) string {
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
