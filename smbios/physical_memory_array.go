// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"encoding/binary"
	"fmt"

	"github.com/digitalocean/go-smbios/smbios"
)

// MemoryArrayLocation represents the memory array location.
type MemoryArrayLocation int

const (
	// MemoryArrayLocationOther is a memory array location type.
	MemoryArrayLocationOther MemoryArrayLocation = iota
	// MemoryArrayLocationUnknown is a memory array location type.
	MemoryArrayLocationUnknown
	// MemoryArrayLocationSystemBoard is a memory array location type.
	MemoryArrayLocationSystemBoard
	// MemoryArrayLocationISAAddon is a memory array location type.
	MemoryArrayLocationISAAddon
	// MemoryArrayLocationEISAAddon is a memory array location type.
	MemoryArrayLocationEISAAddon
	// MemoryArrayLocationPCIAddon is a memory array location type.
	MemoryArrayLocationPCIAddon
	// MemoryArrayLocationMCAAddon is a memory array location type.
	MemoryArrayLocationMCAAddon
	// MemoryArrayLocationPCMCIAAddon is a memory array location type.
	MemoryArrayLocationPCMCIAAddon
	// MemoryArrayLocationProprietaryAddon is a memory array location type.
	MemoryArrayLocationProprietaryAddon
	// MemoryArrayLocationNuBus is a memory array location type.
	MemoryArrayLocationNuBus
	// MemoryArrayLocationPC98C20Addon is a memory array location type.
	MemoryArrayLocationPC98C20Addon
	// MemoryArrayLocationPC98C24Addon is a memory array location type.
	MemoryArrayLocationPC98C24Addon
	// MemoryArrayLocationPC98EAddon is a memory array location type.
	MemoryArrayLocationPC98EAddon
	// MemoryArrayLocationPC98LocalBusAddon is a memory array location type.
	MemoryArrayLocationPC98LocalBusAddon
)

const (
	memoryArrayLocationOther             = "Other"
	memoryArrayLocationUnknown           = "Unknown"
	memoryArrayLocationSystemBoard       = "System board or motherboard"
	memoryArrayLocationISAAddon          = "ISA add-on card"
	memoryArrayLocationEISAAddon         = "EISA add-on card"
	memoryArrayLocationPCIAddon          = "PCI add-on card"
	memoryArrayLocationMCAAddon          = "MCA add-on card"
	memoryArrayLocationPCMCIAAddon       = "PCMCIA add-on card"
	memoryArrayLocationProprietaryAddon  = "Proprietary add-on card "
	memoryArrayLocationNuBus             = "NuBus"
	memoryArrayLocationPC98C20Addon      = "PC-98/C20 add-on card"
	memoryArrayLocationPC98C24Addon      = "PC-98/C24 add-on card"
	memoryArrayLocationPC98EAddon        = "PC-98/E add-on card"
	memoryArrayLocationPC98LocalBusAddon = "PC-98/Local bus add-on card"
)

// String returns the string representation of `MemoryArrayLocation`.
func (m MemoryArrayLocation) String() string {
	return [...]string{
		"", // Placeholder since values start at 01h.
		memoryArrayLocationOther,
		memoryArrayLocationUnknown,
		memoryArrayLocationSystemBoard,
		memoryArrayLocationISAAddon,
		memoryArrayLocationEISAAddon,
		memoryArrayLocationPCIAddon,
		memoryArrayLocationMCAAddon,
		memoryArrayLocationPCMCIAAddon,
		memoryArrayLocationProprietaryAddon,
		memoryArrayLocationNuBus,
		memoryArrayLocationPC98C20Addon,
		memoryArrayLocationPC98C24Addon,
		memoryArrayLocationPC98EAddon,
		memoryArrayLocationPC98LocalBusAddon,
	}[m]
}

// MemoryArrayUse represents memory array use.
type MemoryArrayUse int

const (
	// MemoryArrayUseOther is a memory array use type.
	MemoryArrayUseOther MemoryArrayUse = iota
	// MemoryArrayUseUnknown is a memory array use type.
	MemoryArrayUseUnknown
	// MemoryArrayUseSystemMemory is a memory array use type.
	MemoryArrayUseSystemMemory
	// MemoryArrayUseVideoMemory is a memory array use type.
	MemoryArrayUseVideoMemory
	// MemoryArrayUseFlashMemory is a memory array use type.
	MemoryArrayUseFlashMemory
	// MemoryArrayUseNonVolatileRAM is a memory array use type.
	MemoryArrayUseNonVolatileRAM
	// MemoryArrayUseCacheMemory is a memory array use type.
	MemoryArrayUseCacheMemory
)

const (
	memoryArrayUseOther          = "Other"
	memoryArrayUseUnknown        = "Unknown"
	memoryArrayUseSystemMemory   = "System memory"
	memoryArrayUseVideoMemory    = "Video memory"
	memoryArrayUseFlashMemory    = "Flash memory"
	memoryArrayUseNonVolatileRAM = "Non-volatile RAM"
	memoryArrayUseCacheMemory    = "Cache memory"
)

// String returns the string representation of `MemoryArrayUse`.
func (m MemoryArrayUse) String() string {
	return [...]string{
		"", // Placeholder since values start at 01h.
		memoryArrayUseOther,
		memoryArrayUseUnknown,
		memoryArrayUseSystemMemory,
		memoryArrayUseVideoMemory,
		memoryArrayUseFlashMemory,
		memoryArrayUseNonVolatileRAM,
		memoryArrayUseCacheMemory,
	}[m]
}

// MemoryArrayMemoryErrorCorrection represents memory array memory error correction.
type MemoryArrayMemoryErrorCorrection int

const (
	// MemoryArrayMemoryErrorCorrectionOther a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionOther MemoryArrayMemoryErrorCorrection = iota
	// MemoryArrayMemoryErrorCorrectionUnknown a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionUnknown
	// MemoryArrayMemoryErrorCorrectionNone a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionNone
	// MemoryArrayMemoryErrorCorrectionParity a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionParity
	// MemoryArrayMemoryErrorCorrectionSingleBitECC a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionSingleBitECC
	// MemoryArrayMemoryErrorCorrectionMultiBitECC a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionMultiBitECC
	// MemoryArrayMemoryErrorCorrectionCRC a memory array error correction type.
	MemoryArrayMemoryErrorCorrectionCRC
)

const (
	memoryArrayMemoryErrorCorrectionOther        = "Other"
	memoryArrayMemoryErrorCorrectionUnknown      = "Unknown"
	memoryArrayMemoryErrorCorrectionNone         = "None"
	memoryArrayMemoryErrorCorrectionParity       = "Parity"
	memoryArrayMemoryErrorCorrectionSingleBitECC = "Single-bit ECC"
	memoryArrayMemoryErrorCorrectionMultiBitECC  = "Multi-bit ECC"
	memoryArrayMemoryErrorCorrectionCRC          = "CRC"
)

// String returns the string representation of MemoryArrayMemoryErrorCorrection.
func (m MemoryArrayMemoryErrorCorrection) String() string {
	return [...]string{
		"", // Placeholder since values start at 01h.
		memoryArrayMemoryErrorCorrectionOther,
		memoryArrayMemoryErrorCorrectionUnknown,
		memoryArrayMemoryErrorCorrectionNone,
		memoryArrayMemoryErrorCorrectionParity,
		memoryArrayMemoryErrorCorrectionSingleBitECC,
		memoryArrayMemoryErrorCorrectionMultiBitECC,
		memoryArrayMemoryErrorCorrectionCRC,
	}[m]
}

// PhysicalMemoryArrayStructure represents the SMBIOS physical memory array structure.
type PhysicalMemoryArrayStructure struct {
	*smbios.Structure
}

// PhysicalMemoryArray returns a `PhysicalMemoryArrayStructure`.
func (s *SMBIOS) PhysicalMemoryArray() PhysicalMemoryArrayStructure {
	return s.PhysicalMemoryArrayStructure
}

// Location returns the physical location of the Memory Array,
// whether on the system board or an add-in board.
// See 7.17.1 for definitions.
func (s PhysicalMemoryArrayStructure) Location() MemoryArrayLocation {
	return MemoryArrayLocation(s.Formatted[0])
}

// Use returns the function for which the array is used. See
// 7.17.2 for definitions.
func (s PhysicalMemoryArrayStructure) Use() MemoryArrayUse {
	return MemoryArrayUse(s.Formatted[1])
}

// MemoryErrorCorrection returns the primary hardware error correction or
// detection method supported by this memory array.
// See 7.17.3 for definitions.
func (s PhysicalMemoryArrayStructure) MemoryErrorCorrection() MemoryArrayMemoryErrorCorrection {
	return MemoryArrayMemoryErrorCorrection(s.Formatted[2])
}

// MaximumCapacity represents the physical memory array maximum capacity.
type MaximumCapacity uint32

// String returns the string representation of a `MaximumCapacity`.
func (m MaximumCapacity) String() string {
	if m == 0x80000000 {
		return ""
	}

	n := m / (1024 * 1024)

	return fmt.Sprintf("%d GB", n)
}

// MaximumCapacity returns the maximum memory capacity, in kilobytes, for
// this array. If the capacity is not represented in this
// field, then this field contains 8000 0000h and the
// Extended Maximum Capacity field should be
// used. Values 2 TB (8000 0000h) or greater must
// be represented in the Extended Maximum
// Capacity field.
func (s PhysicalMemoryArrayStructure) MaximumCapacity() MaximumCapacity {
	return MaximumCapacity(binary.LittleEndian.Uint32(s.Formatted[3:7]))
}

// MemoryErrorInformationHandle returns the handle, or instance number, associated with
// any error that was previously detected for the
// array. If the system does not provide the error
// information structure, the field contains FFFEh;
// otherwise, the field contains either FFFFh (if no
// error was detected) or the handle of the error-information structure.
// See 7.18.4 and 7.34.
func (s PhysicalMemoryArrayStructure) MemoryErrorInformationHandle() MemoryErrorInformationHandle {
	return MemoryErrorInformationHandle(binary.LittleEndian.Uint16(s.Formatted[7:9]))
}

// NumberOfMemoryDevices returns the number of slots or sockets available for
// Memory Devices in this array. This value
// represents the number of Memory Device
// structures that comprise this Memory Array. Each
// Memory Device has a reference to the “owning”
// Memory Array.
func (s PhysicalMemoryArrayStructure) NumberOfMemoryDevices() uint16 {
	return binary.LittleEndian.Uint16(s.Formatted[9:11])
}

// ExtendedMaximumCapacity represents the physical memory array extended maximum capacity.
type ExtendedMaximumCapacity uint64

func (m ExtendedMaximumCapacity) String() string {
	if m == 0x00000000 {
		return ""
	}

	n := m / (1024 * 1024)

	return fmt.Sprintf("%d GB", n)
}

// ExtendedMaximumCapacity returns the maximum memory capacity, in bytes, for this
// array. This field is only valid when the Maximum
// Capacity field contains 8000 0000h. When
// Maximum Capacity contains a value that is not
// 8000 0000h, Extended Maximum Capacity must
// contain zeros.
func (s PhysicalMemoryArrayStructure) ExtendedMaximumCapacity() ExtendedMaximumCapacity {
	return ExtendedMaximumCapacity(binary.LittleEndian.Uint64(s.Formatted[11:19]))
}
