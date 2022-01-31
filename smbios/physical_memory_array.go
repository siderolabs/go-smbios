// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"fmt"

	"github.com/digitalocean/go-smbios/smbios"
)

// PhysicalMemoryArray represents the SMBIOS physical memory array.
type PhysicalMemoryArray struct {
	// Location returns the physical location of the Memory Array,
	// whether on the system board or an add-in board.
	// See 7.17.1 for definitions.
	Location MemoryArrayLocation
	// Use returns the function for which the array is used. See
	// 7.17.2 for definitions.
	Use MemoryArrayUse
	// MemoryErrorCorrection returns the primary hardware error correction or
	// detection method supported by this memory array.
	// See 7.17.3 for definitions.
	MemoryErrorCorrection MemoryArrayMemoryErrorCorrection
	// MaximumCapacity returns the maximum memory capacity, in kilobytes, for
	// this array. If the capacity is not represented in this
	// field, then this field contains 8000 0000h and the
	// Extended Maximum Capacity field should be
	// used. Values 2 TB (8000 0000h) or greater must
	// be represented in the Extended Maximum
	// Capacity field.
	MaximumCapacity MaximumCapacity
	// MemoryErrorInformationHandle returns the handle, or instance number, associated with
	// any error that was previously detected for the
	// array. If the system does not provide the error
	// information structure, the field contains FFFEh;
	// otherwise, the field contains either FFFFh (if no
	// error was detected) or the handle of the error-information structure.
	// See 7.18.4 and 7.34.
	MemoryErrorInformationHandle MemoryErrorInformationHandle
	// NumberOfMemoryDevices returns the number of slots or sockets available for
	// Memory Devices in this array. This value
	// represents the number of Memory Device
	// structures that comprise this Memory Array. Each
	// Memory Device has a reference to the “owning”
	// Memory Array.
	NumberOfMemoryDevices uint16
	// ExtendedMaximumCapacity returns the maximum memory capacity, in bytes, for this
	// array. This field is only valid when the Maximum
	// Capacity field contains 8000 0000h. When
	// Maximum Capacity contains a value that is not
	// 8000 0000h, Extended Maximum Capacity must
	// contain zeros.
	ExtendedMaximumCapacity ExtendedMaximumCapacity
}

// NewPhysicalMemoryArray initializes and returns a new `PhysicalMemoryArray`.
func NewPhysicalMemoryArray(s *smbios.Structure) *PhysicalMemoryArray {
	return &PhysicalMemoryArray{
		Location:                     MemoryArrayLocation(GetByte(s, 0x04)),
		Use:                          MemoryArrayUse(GetByte(s, 0x05)),
		MemoryErrorCorrection:        MemoryArrayMemoryErrorCorrection(GetByte(s, 0x06)),
		MaximumCapacity:              MaximumCapacity(GetDWord(s, 0x07)),
		MemoryErrorInformationHandle: MemoryErrorInformationHandle(GetWord(s, 0x0B)),
		NumberOfMemoryDevices:        GetWord(s, 0x0D),
		ExtendedMaximumCapacity:      ExtendedMaximumCapacity(GetQWord(s, 0x0F)),
	}
}

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

// String returns the string representation of `MemoryArrayLocation`.
func (m MemoryArrayLocation) String() string {
	switch m {
	case MemoryArrayLocationOther:
		return _Other
	case MemoryArrayLocationUnknown:
		return _Unknown
	case MemoryArrayLocationSystemBoard:
		return "System board or motherboard"
	case MemoryArrayLocationISAAddon:
		return "ISA add-on card"
	case MemoryArrayLocationEISAAddon:
		return "EISA add-on card"
	case MemoryArrayLocationPCIAddon:
		return "PCI add-on card"
	case MemoryArrayLocationMCAAddon:
		return "MCA add-on card"
	case MemoryArrayLocationPCMCIAAddon:
		return "PCMCIA add-on card"
	case MemoryArrayLocationProprietaryAddon:
		return "Proprietary add-on card "
	case MemoryArrayLocationNuBus:
		return "NuBus"
	case MemoryArrayLocationPC98C20Addon:
		return "PC-98/C20 add-on card"
	case MemoryArrayLocationPC98C24Addon:
		return "PC-98/C24 add-on card"
	case MemoryArrayLocationPC98EAddon:
		return "PC-98/E add-on card"
	case MemoryArrayLocationPC98LocalBusAddon:
		return "PC-98/Local bus add-on card"
	}

	return _Unknown
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

// String returns the string representation of `MemoryArrayUse`.
func (m MemoryArrayUse) String() string {
	switch m {
	case MemoryArrayUseOther:
		return _Other
	case MemoryArrayUseUnknown:
		return _Unknown
	case MemoryArrayUseSystemMemory:
		return "System memory"
	case MemoryArrayUseVideoMemory:
		return "Video memory"
	case MemoryArrayUseFlashMemory:
		return "Flash memory"
	case MemoryArrayUseNonVolatileRAM:
		return "Non-volatile RAM"
	case MemoryArrayUseCacheMemory:
		return "Cache memory"
	}

	return _Unknown
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

// String returns the string representation of MemoryArrayMemoryErrorCorrection.
func (m MemoryArrayMemoryErrorCorrection) String() string {
	switch m {
	case MemoryArrayMemoryErrorCorrectionOther:
		return _Other
	case MemoryArrayMemoryErrorCorrectionUnknown:
		return _Unknown
	case MemoryArrayMemoryErrorCorrectionNone:
		return "None"
	case MemoryArrayMemoryErrorCorrectionParity:
		return "Parity"
	case MemoryArrayMemoryErrorCorrectionSingleBitECC:
		return "Single-bit ECC"
	case MemoryArrayMemoryErrorCorrectionMultiBitECC:
		return "Multi-bit ECC"
	case MemoryArrayMemoryErrorCorrectionCRC:
		return "CRC"
	}

	return _Unknown
}

// MaximumCapacity represents the physical memory
// array maximum memory capacity, in kilobytes.
// If the capacity is not represented in this
// field, then this field contains 8000 0000h and the
// Extended Maximum Capacity field should be
// used. Values 2 TB (8000 0000h) or greater must
// be represented in the Extended Maximum
// Capacity field.
type MaximumCapacity uint32

// String returns the string representation of a `MaximumCapacity`.
func (m MaximumCapacity) String() string {
	if m == 0x80000000 {
		return ""
	}

	n := m / (1024 * 1024)

	return fmt.Sprintf("%d GB", n)
}

// ExtendedMaximumCapacity represents the physical memory
// array extended maximum capacity, in bytes.
// This field is only valid when the Maximum
// Capacity field contains 8000 0000h. When
// Maximum Capacity contains a value that is not
// 8000 0000h, Extended Maximum Capacity must
// contain zeros.
type ExtendedMaximumCapacity uint64

func (m ExtendedMaximumCapacity) String() string {
	if m == 0x00000000 {
		return ""
	}

	n := m / (1024 * 1024)

	return fmt.Sprintf("%d GB", n)
}
