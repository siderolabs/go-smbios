// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"fmt"
	"strconv"

	"github.com/digitalocean/go-smbios/smbios"
)

// MemoryDevice represents a SMBIOS memory device.
//
//nolint:govet
type MemoryDevice struct {
	// PhysicalMemoryArrayHandle is the handle, or instance number, associated with
	// the Physical Memory Array to which this device belongs.
	PhysicalMemoryArrayHandle PhysicalMemoryArrayHandle
	// MemoryErrorInformationHandle is the handle, or instance number, associated with
	// any error that was previously detected for the
	// device. if the system does not provide the error
	// information structure, the field contains FFFEh;
	// otherwise, the field contains either FFFFh (if no
	// error was detected) or the handle of the error information structure.
	// See 7.18.4 and 7.34.
	MemoryErrorInformationHandle MemoryErrorInformationHandle
	// TotalWidth is the total width, in bits, of this memory device,
	// including any check or error-correction bits. If there
	// are no error-correction bits, this value should be
	// equal to Data Width. If the width is unknown, the
	// field is set to FFFFh.
	TotalWidth MemoryDeviceWidth
	// DataWidth is the data width, in bits, of this memory device. A
	// Data Width of 0 and a Total Width of 8 indicates that
	// the device is being used solely to provide 8 error correction bits.
	//  If the width is unknown, the field is set to FFFFh.
	DataWidth MemoryDeviceWidth
	// Size is the size of the memory device. If the value is 0, no
	// memory device is installed in the socket; if the size
	// is unknown, the field value is FFFFh. If the size is
	// 32 GB-1 MB or greater, the field value is 7FFFh and
	// the actual size is stored in the Extended Size field.
	// The granularity in which the value is specified
	// depends on the setting of the most-significant bit (bit
	// 15). If the bit is 0, the value is specified in megabyte
	// units; if the bit is 1, the value is specified in kilobyte
	// units. For example, the value 8100h identifies a
	// 256 KB memory device and 0100h identifies a
	// 256 MB memory device.
	Size MemoryDeviceSize
	// FormFactor is the implementation form factor for this memory
	// device. See 7.18.1 for definitions.
	FormFactor FormFactor
	// DeviceSet identifies when the Memory Device is one of a set
	// of Memory Devices that must be populated with all
	// devices of the same type and size, and the set to
	// which this device belongs. A value of 0 indicates
	// that the device is not part of a set; a value of FFh
	// indicates that the attribute is unknown.
	// NOTE: A Device Set number must be unique within the
	// context of the Memory Array containing this Memory
	// Device.
	DeviceSet string
	// DeviceLocator is the string number of the string that identifies the
	// physically-labeled socket or board position where
	// the memory device is located
	// EXAMPLE: “SIMM 3”.
	DeviceLocator string
	// BankLocator is the string number of the string that identifies the
	// physically labeled bank where the memory device is
	// located,
	// EXAMPLE: “Bank 0” or “A”.
	BankLocator string
	// MemoryType is the type of memory used in this device; see 7.18.2
	// for definitions.
	MemoryType MemoryType
	// TypeDetail is additional detail on the memory device type; see
	// 7.18.3 for definitions.
	TypeDetail TypeDetail
	// Speed identifies the maximum capable speed of the
	// device, in megatransfers per second (MT/s). See
	// 7.18.4 for details.
	// 0000h = the speed is unknown
	// FFFFh = the speed is 65,535 MT/s or greater,
	// and the actual speed is stored in the Extended
	// Speed field.
	Speed MemoryDeviceSpeed
	// Manufacturer is the string number for the manufacturer of this memory
	// device.
	Manufacturer string
	// SerialNumber is the string number for the serial number of this memory
	// device.
	// This value is set by the manufacturer and normally
	// is not changeable.
	SerialNumber string
	// AssetTag is the string number for the asset tag of this memory
	// device.
	AssetTag string
	// PartNumber is the string number for the part number of this memory
	// device.
	// This value is set by the manufacturer and normally
	// is not changeable.
	PartNumber string
	// Attributes returns the memory device attributes.
	Attributes uint8
	// ExtendedSize returns the memory device extended size.
	// The Extended Size field is intended to represent memory devices larger than 32,767 MB (32 GB - 1 MB),
	// which cannot be described using the Size field. This field is only meaningful if the value in the Size field
	// is 7FFFh. For compatibility with older SMBIOS parsers, memory devices smaller than (32 GB - 1 MB)
	// should be represented using their size in the Size field, leaving the Extended Size field set to 0.
	// Bit 31 is reserved for future use and must be set to 0.
	// Bits 30:0 represent the size of the memory device in megabytes.
	// Example: 0000_8000h indicates a 32 GB memory device (32,768 MB), 0002_0000h represent
	ExtendedSize MemoryDeviceExtendedSize
	// ConfiguredMemorySpeed identifies the configured speed of the memory
	// device, in megatransfers per second (MT/s). See
	// 7.18.4 for details.
	// 0000h = the speed is unknown
	// FFFFh = the speed is 65,535 MT/s or greater,
	// and the actual speed is stored in the Extended
	// Configured Memory Speed field.
	ConfiguredMemorySpeed MemoryDeviceSpeed
	// MinimumVoltage is the minimum operating voltage for this device, in millivolts
	// If the value is 0, the voltage is unknown.
	MinimumVoltage MemoryDeviceVoltage
	// MaximumVoltage is the maximum operating voltage for this device, in millivolts
	// If the value is 0, the voltage is unknown.
	MaximumVoltage MemoryDeviceVoltage
	// ConfiguredVoltage is the configured voltage for this device, in millivolts
	// If the value is 0, the voltage is unknown.
	ConfiguredVoltage MemoryDeviceVoltage
}

// NewMemoryDevice initializes and returns a new `MemoryDevice`.
func NewMemoryDevice(s *smbios.Structure) *MemoryDevice {
	return &MemoryDevice{
		PhysicalMemoryArrayHandle:    PhysicalMemoryArrayHandle(GetWord(s, 0x04)),
		MemoryErrorInformationHandle: MemoryErrorInformationHandle(GetWord(s, 0x06)),
		TotalWidth:                   MemoryDeviceWidth(GetWord(s, 0x08)),
		DataWidth:                    MemoryDeviceWidth(GetWord(s, 0x0A)),
		Size:                         MemoryDeviceSize(GetWord(s, 0x0C)),
		FormFactor:                   FormFactor(GetByte(s, 0x0E)),
		DeviceSet:                    _GetDeviceSet(s, 0x0F),
		DeviceLocator:                GetStringOrEmpty(s, 0x10),
		BankLocator:                  GetStringOrEmpty(s, 0x11),
		MemoryType:                   MemoryType(GetByte(s, 0x12)),
		TypeDetail:                   TypeDetail(GetByte(s, 0x13)),
		Speed:                        MemoryDeviceSpeed(GetWord(s, 0x15)),
		Manufacturer:                 GetStringOrEmpty(s, 0x17),
		SerialNumber:                 GetStringOrEmpty(s, 0x18),
		AssetTag:                     GetStringOrEmpty(s, 0x19),
		PartNumber:                   GetStringOrEmpty(s, 0x1A),
		Attributes:                   GetByte(s, 0x1B),
		ExtendedSize:                 MemoryDeviceExtendedSize(GetDWord(s, 0x1C)),
		ConfiguredMemorySpeed:        MemoryDeviceSpeed(GetWord(s, 0x20)),
		MinimumVoltage:               MemoryDeviceVoltage(GetWord(s, 0x22)),
		MaximumVoltage:               MemoryDeviceVoltage(GetWord(s, 0x24)),
		ConfiguredVoltage:            MemoryDeviceVoltage(GetWord(s, 0x26)),
	}
}

// PhysicalMemoryArrayHandle represents the SMBIOS physical memory array handle.
type PhysicalMemoryArrayHandle uint16

func (p PhysicalMemoryArrayHandle) String() string {
	return fmt.Sprintf("0x%X", uint16(p))
}

// MemoryErrorInformationHandle represents the SMBIOS memory error information handle.
type MemoryErrorInformationHandle uint16

// String returns the string representation of the SMBIOS memory error information handle.
func (m MemoryErrorInformationHandle) String() string {
	if m == 0xFFFE {
		return "Not Provided"
	}

	if m == 0xFFFF {
		return "No Error Detected"
	}

	return fmt.Sprintf("0x%X", uint16(m))
}

// MemoryDeviceWidth represents the SMBIOS memory device width.
type MemoryDeviceWidth uint16

// String returns the string representation of the SMBIOS memory device width.
func (m MemoryDeviceWidth) String() string {
	if m == 0xFFFF {
		return _Unknown
	}

	return fmt.Sprintf("%d bits", m)
}

// MemoryDeviceSize represents the SMBIOS memory device size.
type MemoryDeviceSize uint16

// Megabytes returns the size of the SMBIOS memory device converted to megabytes.
func (m MemoryDeviceSize) Megabytes() int {
	if IsNthBitSet(int(m), 15) {
		return int(m) / 1024
	}

	return int(m)
}

// String returns the string representation of the SMBIOS memory device size.
func (m MemoryDeviceSize) String() string {
	if m == 0xFFFF {
		return _Unknown
	}

	if m == 0x7FFF {
		return "See Extended Size"
	}

	var units string
	if IsNthBitSet(int(m), 15) {
		units = "KB"
	} else {
		units = "MB"
	}

	return fmt.Sprintf("%d %s", m, units)
}

// MemoryDeviceExtendedSize represents the SMBIOS memory device extended size.
type MemoryDeviceExtendedSize uint16

// String returns the string representation of the SMBIOS memory device extended size.
func (m MemoryDeviceExtendedSize) String() string {
	return fmt.Sprintf("%d MB", m)
}

// FormFactor represents the SMBIOS memory device form factor.
type FormFactor int

const (
	// FormFactorOther is a memory device form factor type.
	FormFactorOther FormFactor = iota
	// FormFactorUnknown is a memory device form factor type.
	FormFactorUnknown
	// FormFactorSIMM is a memory device form factor type.
	FormFactorSIMM
	// FormFactorSIP is a memory device form factor type.
	FormFactorSIP
	// FormFactorChip is a memory device form factor type.
	FormFactorChip
	// FormFactorDIP is a memory device form factor type.
	FormFactorDIP
	// FormFactorZIP is a memory device form factor type.
	FormFactorZIP
	// FormFactorProprietaryCard is a memory device form factor type.
	FormFactorProprietaryCard
	// FormFactorDIMM is a memory device form factor type.
	FormFactorDIMM
	// FormFactorTSOP is a memory device form factor type.
	FormFactorTSOP
	// FormFactorRowOfChips is a memory device form factor type.
	FormFactorRowOfChips
	// FormFactorRIMM is a memory device form factor type.
	FormFactorRIMM
	// FormFactorSODIMM is a memory device form factor type.
	FormFactorSODIMM
	// FormFactorSRIMM is a memory device form factor type.
	FormFactorSRIMM
	// FormFactorFBDIMM is a memory device form factor type.
	FormFactorFBDIMM
	// FormFactorFBDie is a memory device form factor type.
	FormFactorFBDie
)

// String returns the string representation of a form factor.
func (f FormFactor) String() string {
	switch f {
	case FormFactorOther:
		return _Other
	case FormFactorUnknown:
		return _Unknown
	case FormFactorSIMM:
		return "SIMM"
	case FormFactorSIP:
		return "SIP"
	case FormFactorChip:
		return "Chip"
	case FormFactorDIP:
		return "DIP"
	case FormFactorZIP:
		return "ZIP"
	case FormFactorProprietaryCard:
		return "Proprietary Card"
	case FormFactorDIMM:
		return "DIMM"
	case FormFactorTSOP:
		return "TSOP"
	case FormFactorRowOfChips:
		return "Row of chips"
	case FormFactorRIMM:
		return "RIMM"
	case FormFactorSODIMM:
		return "SODIMM"
	case FormFactorSRIMM:
		return "SRIMM"
	case FormFactorFBDIMM:
		return "FB-DIMM"
	case FormFactorFBDie:
		return "Die"
	}

	return _Unknown
}

func _GetDeviceSet(s *smbios.Structure, offset int) string {
	b := GetByte(s, offset)

	if b == 0 {
		return _Empty
	}

	if b == 0xFF {
		return _Empty
	}

	return strconv.FormatInt(int64(b), 10)
}

// MemoryType represents the SMBIOS memory device type.
type MemoryType int

const (
	// MemoryTypeOther is memory device type.
	MemoryTypeOther MemoryType = iota
	// MemoryTypeUnknown is memory device type.
	MemoryTypeUnknown
	// MemoryTypeDRAM is memory device type.
	MemoryTypeDRAM
	// MemoryTypeEDRAM is memory device type.
	MemoryTypeEDRAM
	// MemoryTypeVRAM is memory device type.
	MemoryTypeVRAM
	// MemoryTypeSRAM is memory device type.
	MemoryTypeSRAM
	// MemoryTypeRAM is memory device type.
	MemoryTypeRAM
	// MemoryTypeROM is memory device type.
	MemoryTypeROM
	// MemoryTypeFLASH is memory device type.
	MemoryTypeFLASH
	// MemoryTypeEEPROM is memory device type.
	MemoryTypeEEPROM
	// MemoryTypeFEPROM is memory device type.
	MemoryTypeFEPROM
	// MemoryTypeEPROM is memory device type.
	MemoryTypeEPROM
	// MemoryTypeCDRAM is memory device type.
	MemoryTypeCDRAM
	// MemoryType3DRAM is memory device type.
	MemoryType3DRAM
	// MemoryTypeSDRAM is memory device type.
	MemoryTypeSDRAM
	// MemoryTypeSGRAM is memory device type.
	MemoryTypeSGRAM
	// MemoryTypeRDRAM is memory device type.
	MemoryTypeRDRAM
	// MemoryTypeDDR is memory device type.
	MemoryTypeDDR
	// MemoryTypeDDR2 is memory device type.
	MemoryTypeDDR2
	// MemoryTypeDDR2FBDIMM is memory device type.
	MemoryTypeDDR2FBDIMM
	// MemoryTypeReserved is memory device type.
	MemoryTypeReserved
	// MemoryTypeDDR3 is memory device type.
	MemoryTypeDDR3
	// MemoryTypeFBD2 is memory device type.
	MemoryTypeFBD2
	// MemoryTypeDDR4 is memory device type.
	MemoryTypeDDR4
	// MemoryTypeLPDDR is memory device type.
	MemoryTypeLPDDR
	// MemoryTypeLPDDR2 is memory device type.
	MemoryTypeLPDDR2
	// MemoryTypeLPDDR3 is memory device type.
	MemoryTypeLPDDR3
	// MemoryTypeLPDDR4 is memory device type.
	MemoryTypeLPDDR4
	// MemoryTypeLogicalNonVolatileDevice is memory device type.
	MemoryTypeLogicalNonVolatileDevice
	// MemoryTypeHBM is memory device type.
	MemoryTypeHBM
	// MemoryTypeHBM2 is memory device type.
	MemoryTypeHBM2
	// MemoryTypeDDR5 is memory device type.
	MemoryTypeDDR5
	// MemoryTypeLPDDR5 is memory device type.
	MemoryTypeLPDDR5
)

// String returns the string representation of this `MemoryType` enum constant.
//nolint: gocyclo, cyclop
func (m MemoryType) String() string {
	switch m {
	case MemoryTypeOther:
		return _Other
	case MemoryTypeUnknown:
		return _Unknown
	case MemoryTypeDRAM:
		return "DRAM"
	case MemoryTypeEDRAM:
		return "EDRAM"
	case MemoryTypeVRAM:
		return "VRAM"
	case MemoryTypeSRAM:
		return "SRAM"
	case MemoryTypeRAM:
		return "RAM"
	case MemoryTypeROM:
		return "ROM"
	case MemoryTypeFLASH:
		return "FLASH"
	case MemoryTypeEEPROM:
		return "EEPROM"
	case MemoryTypeFEPROM:
		return "FEPROM"
	case MemoryTypeEPROM:
		return "EPROM"
	case MemoryTypeCDRAM:
		return "CDRAM"
	case MemoryType3DRAM:
		return "3DRAM"
	case MemoryTypeSDRAM:
		return "SDRAM"
	case MemoryTypeSGRAM:
		return "SGRAM"
	case MemoryTypeRDRAM:
		return "RDRAM"
	case MemoryTypeDDR:
		return "DDR"
	case MemoryTypeDDR2:
		return "DDR2"
	case MemoryTypeDDR2FBDIMM:
		return "DDR2 FB-DIMM"
	case MemoryTypeReserved:
		return _Reserved
	case MemoryTypeDDR3:
		return "DDR3"
	case MemoryTypeFBD2:
		return "FBD2"
	case MemoryTypeDDR4:
		return "DDR4"
	case MemoryTypeLPDDR:
		return "LPDDR"
	case MemoryTypeLPDDR2:
		return "LPDDR2"
	case MemoryTypeLPDDR3:
		return "LPDDR3"
	case MemoryTypeLPDDR4:
		return "LPDDR4"
	case MemoryTypeLogicalNonVolatileDevice:
		return "Logical non-volatile device"
	case MemoryTypeHBM:
		return "HBM (High Bandwidth Memory)"
	case MemoryTypeHBM2:
		return "HBM2 (High Bandwidth Memory Generation 2)"
	case MemoryTypeDDR5:
		return "DDR5"
	case MemoryTypeLPDDR5:
		return "LPDDR5"
	}

	return _Unknown
}

// TypeDetail represents the SMBIOS memory device detail.
type TypeDetail uint16

// TypeDetailAttribute represents the SMBIOS memory device detail attribute.
type TypeDetailAttribute int

const (
	// TypeDetailAttributeLRDIMM is a memory device type detail attribute.
	TypeDetailAttributeLRDIMM = iota
	// TypeDetailAttributeUnbuffered is a memory device type detail attribute.
	TypeDetailAttributeUnbuffered
	// TypeDetailAttributeRegistered is a memory device type detail attribute.
	TypeDetailAttributeRegistered
	// TypeDetailAttributeNonVolatile is a memory device type detail attribute.
	TypeDetailAttributeNonVolatile
	// TypeDetailAttributeCacheDRAM is a memory device type detail attribute.
	TypeDetailAttributeCacheDRAM
	// TypeDetailAttributeWindowDRAM is a memory device type detail attribute.
	TypeDetailAttributeWindowDRAM
	// TypeDetailAttributeEDO is a memory device type detail attribute.
	TypeDetailAttributeEDO
	// TypeDetailAttributeCMOS is a memory device type detail attribute.
	TypeDetailAttributeCMOS
	// TypeDetailAttributeSynchronous is a memory device type detail attribute.
	TypeDetailAttributeSynchronous
	// TypeDetailAttributeRAMBUS is a memory device type detail attribute.
	TypeDetailAttributeRAMBUS
	// TypeDetailAttributePseudoStatic is a memory device type detail attribute.
	TypeDetailAttributePseudoStatic
	// TypeDetailAttributeStaticColumn is a memory device type detail attribute.
	TypeDetailAttributeStaticColumn
	// TypeDetailAttributeFastPaged is a memory device type detail attribute.
	TypeDetailAttributeFastPaged
	// TypeDetailAttributeUnknown is a memory device type detail attribute.
	TypeDetailAttributeUnknown
	// TypeDetailAttributeOther is a memory device type detail attribute.
	TypeDetailAttributeOther
	// TypeDetailAttributeReserved is a memory device type detail attribute.
	TypeDetailAttributeReserved
)

func (t TypeDetailAttribute) String() string {
	switch t {
	case TypeDetailAttributeLRDIMM:
		return "LRDIMM" // Bit 15
	case TypeDetailAttributeUnbuffered:
		return "Unbuffered (Unregistered)" // Bit 14
	case TypeDetailAttributeRegistered:
		return "Registered (Buffered)" // Bit 13
	case TypeDetailAttributeNonVolatile:
		return "Non-volatile" // Bit 12
	case TypeDetailAttributeCacheDRAM:
		return "Cache DRAM" // Bit 11
	case TypeDetailAttributeWindowDRAM:
		return "Window DRAM" // Bit 10
	case TypeDetailAttributeEDO:
		return "EDO" // Bit 9
	case TypeDetailAttributeCMOS:
		return "CMOS" // Bit 8
	case TypeDetailAttributeSynchronous:
		return "Synchronous" // Bit 7
	case TypeDetailAttributeRAMBUS:
		return "RAMBUS" // Bit 6
	case TypeDetailAttributePseudoStatic:
		return "Pseudo-static" // Bit 5
	case TypeDetailAttributeStaticColumn:
		return "Static column" // Bit 4
	case TypeDetailAttributeFastPaged:
		return "Fast-paged" // Bit 3
	case TypeDetailAttributeUnknown:
		return _Unknown // Bit 2
	case TypeDetailAttributeOther:
		return _Other // Bit 1
	case TypeDetailAttributeReserved:
		return _Reserved // Bit 0
	}

	return _Unknown
}

// Attributes returns the SMBIOS memory device detail attribute.
func (t TypeDetail) Attributes() []TypeDetailAttribute {
	b := bits(int(t))

	attributes := []TypeDetailAttribute{}

	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == 1 {
			attributes = append(attributes, TypeDetailAttribute(i))
		}
	}

	return attributes
}

func (t TypeDetail) String() string {
	attributes := t.Attributes()

	s := ""

	for _, attribute := range attributes {
		if s == "" {
			s += attribute.String()

			continue
		}

		s += " " + attribute.String()
	}

	return s
}

// MemoryDeviceSpeed represents the SMBIOS memory device speed.
type MemoryDeviceSpeed uint16

func (m MemoryDeviceSpeed) String() string {
	return fmt.Sprintf("%d MT/s", m)
}

// MemoryDeviceVoltage represents the SMBIOS memory device voltage.
type MemoryDeviceVoltage uint16

func (m MemoryDeviceVoltage) String() string {
	if m == 0 {
		return _Unknown
	}

	return fmt.Sprintf("%.2f V", float32(m)/1000)
}

func bits(input int) []int {
	b := []int{0}

	n := 1

	for n*2 <= input {
		n *= 2
	}

	for n >= 1 {
		b = append(b, input/n)
		input %= n
		n /= 2
	}

	return b
}
