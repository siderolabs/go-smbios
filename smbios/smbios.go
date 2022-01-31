// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/digitalocean/go-smbios/smbios"
)

// Version represents the System Management BIOS version.
type Version struct {
	Major    int
	Minor    int
	Revision int
}

// SMBIOS represents the System Management BIOS.
type SMBIOS struct { //nolint:govet
	Version    Version
	Structures []*smbios.Structure `json:"-"`

	BIOSInformation            BIOSInformation
	SystemInformation          SystemInformation
	BaseboardInformation       BaseboardInformation
	SystemEnclosure            SystemEnclosure
	ProcessorInformation       []ProcessorInformation
	CacheInformation           []CacheInformation
	PortConnectorInformation   []PortConnectorInformation
	SystemSlots                []SystemSlot
	OEMStrings                 OEMStrings
	SystemConfigurationOptions SystemConfigurationOptions
	BIOSLanguageInformation    BIOSLanguageInformation
	GroupAssociations          GroupAssociations
	PhysicalMemoryArray        PhysicalMemoryArray
	MemoryDevices              []MemoryDevice
}

// New initializes and returns a new `SMBIOS`.
func New() (*SMBIOS, error) {
	rc, ep, err := smbios.Stream()
	if err != nil {
		return nil, fmt.Errorf("failed to open stream: %w", err)
	}

	//nolint: errcheck
	defer rc.Close()

	var version Version
	version.Major, version.Minor, version.Revision = ep.Version()

	return Decode(rc, version)
}

// Decode decodes the stream of the provided `Reader` and returns a new `SMBIOS`.
func Decode(rc io.Reader, version Version) (*SMBIOS, error) {
	s := &SMBIOS{}

	s.Version = version

	d := smbios.NewDecoder(rc)

	structures, err := d.Decode()
	if err != nil {
		return nil, fmt.Errorf("failed to decode structures: %w", err)
	}

	s.Structures = structures
	s._Destructure(structures)

	return s, nil
}

// _Destructure destructures the slice of `Structure`s and
// stores the resulting information inside this `SMBIOS`.
func (s *SMBIOS) _Destructure(structures []*smbios.Structure) {
	for _, structure := range structures {
		switch structure.Header.Type {
		case 0:
			s.BIOSInformation = *NewBIOSInformation(structure)
		case 1:
			s.SystemInformation = *NewSystemInformation(structure, s.Version)
		case 2:
			s.BaseboardInformation = *NewBaseboardInformation(structure)
		case 3:
			s.SystemEnclosure = *NewSystemEnclosure(structure)
		case 4:
			processorInformation := *NewProcessorInformation(structure)
			s.ProcessorInformation = append(s.ProcessorInformation, processorInformation)
		case 5:
			// Obsolete.
		case 6:
			// Obsolete.
		case 7:
			cacheInformation := *NewCacheInformation(structure)
			s.CacheInformation = append(s.CacheInformation, cacheInformation)
		case 8:
			portConnectorInformation := *NewPortConnectorInformation(structure)
			s.PortConnectorInformation = append(s.PortConnectorInformation, portConnectorInformation)
		case 9:
			systemSlot := *NewSystemSlot(structure)
			s.SystemSlots = append(s.SystemSlots, systemSlot)
		case 10:
			// Obsolete.
		case 11:
			s.OEMStrings = *NewOEMStrings(structure)
		case 12:
			s.SystemConfigurationOptions = *NewSystemConfigurationOptions(structure)
		case 13:
			s.BIOSLanguageInformation = *NewBIOSLanguageInformation(structure)
		case 14:
			s.GroupAssociations = *NewGroupAssociations(structure)
		case 15:
			// Unimplemented.
		case 16:
			s.PhysicalMemoryArray = *NewPhysicalMemoryArray(structure)
		case 17:
			memoryDevice := *NewMemoryDevice(structure)
			s.MemoryDevices = append(s.MemoryDevices, memoryDevice)
		}
	}
}

var (
	_Empty    = ""
	_Unknown  = "Unknown"
	_Reserved = "Reserved"
	_Other    = "Other"
)

// GetStrings retrieves all strings in the given structure.
func GetStrings(s *smbios.Structure) []string {
	if s.Strings == nil {
		return []string{}
	}

	return s.Strings
}

// GetStringOrEmpty retrieves a string at the given offset.
// Returns an empty string if no string was present.
func GetStringOrEmpty(s *smbios.Structure, offset int) string {
	index := GetByte(s, offset)

	if index == 0 || int(index) > len(s.Strings) {
		return _Empty
	}

	str := s.Strings[index-1]
	trimmed := strings.TrimSpace(str)

	// Convert to lowercase to address multiple formats:
	//   - "To Be Filled By O.E.M."
	//   - "To be filled by O.E.M."
	if strings.ToLower(trimmed) == "to be filled by o.e.m." {
		return _Empty
	}

	return trimmed
}

// GetByte retrieves a 8-bit unsigned integer at the given offset.
func GetByte(s *smbios.Structure, offset int) uint8 {
	// the `Formatted` byte slice is missing the first 4 bytes of the structure that are stripped out as header info.
	// so we need to subtract 4 from the offset mentioned in the SMBIOS documentation to get the right value.
	index := offset - 4
	if index >= len(s.Formatted) {
		return 0
	}

	return s.Formatted[index]
}

// GetWord retrieves a 16-bit unsigned integer at the given offset.
func GetWord(s *smbios.Structure, offset int) uint16 {
	// the `Formatted` byte slice is missing the first 4 bytes of the structure that are stripped out as header info.
	// so we need to subtract 4 from the offset mentioned in the SMBIOS documentation to get the right value.
	index := offset - 4
	if index >= len(s.Formatted) {
		return 0
	}

	b := s.Formatted[index : index+2]
	if len(b) != 2 {
		return 0
	}

	return binary.LittleEndian.Uint16(b)
}

// GetDWord retrieves a 32-bit unsigned integer at the given offset.
func GetDWord(s *smbios.Structure, offset int) uint32 {
	// the `Formatted` byte slice is missing the first 4 bytes of the structure that are stripped out as header info.
	// so we need to subtract 4 from the offset mentioned in the SMBIOS documentation to get the right value.
	index := offset - 4
	if index >= len(s.Formatted) {
		return 0
	}

	b := s.Formatted[index : index+4]
	if len(b) != 4 {
		return 0
	}

	return binary.LittleEndian.Uint32(b)
}

// GetQWord retrieves a 64-bit unsigned integer at the given offset.
func GetQWord(s *smbios.Structure, offset int) uint64 {
	// the `Formatted` byte slice is missing the first 4 bytes of the structure that are stripped out as header info.
	// so we need to subtract 4 from the offset mentioned in the SMBIOS documentation to get the right value.
	index := offset - 4
	if index >= len(s.Formatted) {
		return 0
	}

	b := s.Formatted[index : index+8]
	if len(b) != 8 {
		return 0
	}

	return binary.LittleEndian.Uint64(b)
}

// IsNthBitSet returns true if the `n`th bit is 1 inside `b`.
func IsNthBitSet(b, n int) bool {
	return b&(1<<n) != 0
}
