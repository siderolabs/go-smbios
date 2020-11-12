// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/digitalocean/go-smbios/smbios"
)

// MemoryDeviceStructure represents the SMBIOS memory device structure.
type MemoryDeviceStructure struct {
	smbios.Structure
}

// MemoryDevice returns a `MemoryDeviceStructure`.
func (s SMBIOS) MemoryDevice() MemoryDeviceStructure {
	return s.MemoryDeviceStructure
}

// PhysicalMemoryArrayHandle represents the SMBIOS physical memory array handle.
type PhysicalMemoryArrayHandle uint16

func (p PhysicalMemoryArrayHandle) String() string {
	return fmt.Sprintf("0x%X", uint16(p))
}

// PhysicalMemoryArrayHandle is the handle, or instance number, associated with
// the Physical Memory Array to which this device belongs.
func (s MemoryDeviceStructure) PhysicalMemoryArrayHandle() PhysicalMemoryArrayHandle {
	return PhysicalMemoryArrayHandle(binary.LittleEndian.Uint16(s.Formatted[0:2]))
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

// MemoryErrorInformationHandle is the handle, or instance number, associated with
// any error that was previously detected for the
// device. if the system does not provide the error
// information structure, the field contains FFFEh;
// otherwise, the field contains either FFFFh (if no
// error was detected) or the handle of the errorinformation structure.
// See 7.18.4 and 7.34.
func (s MemoryDeviceStructure) MemoryErrorInformationHandle() MemoryErrorInformationHandle {
	return MemoryErrorInformationHandle(binary.LittleEndian.Uint16(s.Formatted[2:4]))
}

// MemoryDeviceWidth represents the SMBIOS memory device width.
type MemoryDeviceWidth uint16

// String returns the string representation of the SMBIOS memory device width.
func (m MemoryDeviceWidth) String() string {
	if m == 0xFFFF {
		return typeUnknown
	}

	return fmt.Sprintf("%d bits", m)
}

// TotalWidth is the total width, in bits, of this memory device,
// including any check or error-correction bits. If there
// are no error-correction bits, this value should be
// equal to Data Width. If the width is unknown, the
// field is set to FFFFh.
func (s MemoryDeviceStructure) TotalWidth() MemoryDeviceWidth {
	return MemoryDeviceWidth(binary.LittleEndian.Uint16(s.Formatted[4:6]))
}

// DataWidth is the data width, in bits, of this memory device. A
// Data Width of 0 and a Total Width of 8 indicates that
// the device is being used solely to provide 8 errorcorrection bits.
//  If the width is unknown, the field is set to FFFFh.
func (s MemoryDeviceStructure) DataWidth() MemoryDeviceWidth {
	return MemoryDeviceWidth(binary.LittleEndian.Uint16(s.Formatted[6:8]))
}

// MemoryDeviceSize represents the SMBIOS memory device size.
type MemoryDeviceSize uint16

// String returns the string representation fo the SMBIOS memory device size.
func (m MemoryDeviceSize) String() string {
	if m == 0xFFFF {
		return typeUnknown
	}

	if m == 0x7FFF {
		return "See Extended Size"
	}

	b := bits(int(m))

	units := ""

	if b[15] == 0 {
		units = "MB"
	} else if b[15] == 1 {
		units = "KB"
	}

	return fmt.Sprintf("%d %s", m, units)
}

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
func (s MemoryDeviceStructure) Size() MemoryDeviceSize {
	return MemoryDeviceSize(binary.LittleEndian.Uint16(s.Formatted[8:10]))
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

const (
	formFactorOther           = "Other"
	formFactorUnknown         = "Unknown"
	formFactorSIMM            = "SIMM"
	formFactorSIP             = "SIP"
	formFactorChip            = "Chip"
	formFactorDIP             = "DIP"
	formFactorZIP             = "ZIP"
	formFactorProprietaryCard = "Proprietary Card"
	formFactorDIMM            = "DIMM"
	formFactorTSOP            = "TSOP"
	formFactorRowOfChips      = "Row of chips"
	formFactorRIMM            = "RIMM"
	formFactorSODIMM          = "SODIMM"
	formFactorSRIMM           = "SRIMM"
	formFactorFBDIMM          = "FB-DIMM"
	formFactorDie             = "Die"
)

// String returns the string representation of a form factor.
func (f FormFactor) String() string {
	return [...]string{
		"", // Placeholder since values start at 01h.
		formFactorOther,
		formFactorUnknown,
		formFactorSIMM,
		formFactorSIP,
		formFactorChip,
		formFactorDIP,
		formFactorZIP,
		formFactorProprietaryCard,
		formFactorDIMM,
		formFactorTSOP,
		formFactorRowOfChips,
		formFactorRIMM,
		formFactorSODIMM,
		formFactorSRIMM,
		formFactorFBDIMM,
		formFactorDie,
	}[f]
}

// FormFactor is the implementation form factor for this memory
// device. See 7.18.1 for definitions.
func (s MemoryDeviceStructure) FormFactor() FormFactor {
	return FormFactor(s.Formatted[10])
}

// DeviceSet identifies when the Memory Device is one of a set
// of Memory Devices that must be populated with all
// devices of the same type and size, and the set to
// which this device belongs. A value of 0 indicates
// that the device is not part of a set; a value of FFh
// indicates that the attribute is unknown.
// NOTE: A Device Set number must be unique within the
// context of the Memory Array containing this Memory
// Device.
func (s MemoryDeviceStructure) DeviceSet() string {
	b := s.Formatted[11]

	if b == 0 {
		return "None"
	}

	if b == 0xFF {
		return typeUnknown
	}

	return strconv.FormatInt(int64(b), 10)
}

// Locator is the string number of the string that identifies the
// physically-labeled socket or board position where
// the memory device is located
// EXAMPLE: “SIMM 3”.
func (s MemoryDeviceStructure) Locator() string {
	return get(s.Structure, 0)
}

// BankLocator is the string number of the string that identifies the
// physically labeled bank where the memory device is
// located,
// EXAMPLE: “Bank 0” or “A”.
func (s MemoryDeviceStructure) BankLocator() string {
	return get(s.Structure, 1)
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

const (
	memoryTypeOther                    = "Other"
	memoryTypeUnknown                  = "Unknown"
	memoryTypeDRAM                     = "DRAM"
	memoryTypeEDRAM                    = "EDRAM"
	memoryTypeVRAM                     = "VRAM"
	memoryTypeSRAM                     = "SRAM"
	memoryTypeRAM                      = "RAM"
	memoryTypeROM                      = "ROM"
	memoryTypeFLASH                    = "FLASH"
	memoryTypeEEPROM                   = "EEPROM"
	memoryTypeFEPROM                   = "FEPROM"
	memoryTypeEPROM                    = "EPROM"
	memoryTypeCDRAM                    = "CDRAM"
	memoryType3DRAM                    = "3DRAM"
	memoryTypeSDRAM                    = "SDRAM"
	memoryTypeSGRAM                    = "SGRAM"
	memoryTypeRDRAM                    = "RDRAM"
	memoryTypeDDR                      = "DDR"
	memoryTypeDDR2                     = "DDR2"
	memoryTypeDDR2FBDIMM               = "DDR2 FB-DIMM"
	memoryTypeReserved                 = "Reserved"
	memoryTypeDDR3                     = "DDR3"
	memoryTypeFBD2                     = "FBD2"
	memoryTypeDDR4                     = "DDR4"
	memoryTypeLPDDR                    = "LPDDR"
	memoryTypeLPDDR2                   = "LPDDR2"
	memoryTypeLPDDR3                   = "LPDDR3"
	memoryTypeLPDDR4                   = "LPDDR4"
	memoryTypeLogicalNonVolatileDevice = "Logical non-volatile device"
	memoryTypeHBM                      = "HBM (High Bandwidth Memory)"
	memoryTypeHBM2                     = "HBM2 (High Bandwidth Memory Generation 2)"
	memoryTypeDDR5                     = "DDR5"
	memoryTypeLPDDR5                   = "LPDDR5"
)

func (m MemoryType) String() string {
	return [...]string{
		"", // Placeholder since values start at 01h.
		memoryTypeOther,
		memoryTypeUnknown,
		memoryTypeDRAM,
		memoryTypeEDRAM,
		memoryTypeVRAM,
		memoryTypeSRAM,
		memoryTypeRAM,
		memoryTypeROM,
		memoryTypeFLASH,
		memoryTypeEEPROM,
		memoryTypeFEPROM,
		memoryTypeEPROM,
		memoryTypeCDRAM,
		memoryType3DRAM,
		memoryTypeSDRAM,
		memoryTypeSGRAM,
		memoryTypeRDRAM,
		memoryTypeDDR,
		memoryTypeDDR2,
		memoryTypeDDR2FBDIMM,
		memoryTypeReserved, // First reserved byte.
		memoryTypeReserved, // Second reserved byte.
		memoryTypeReserved, // Third reserved byte.
		memoryTypeDDR3,
		memoryTypeFBD2,
		memoryTypeDDR4,
		memoryTypeLPDDR,
		memoryTypeLPDDR2,
		memoryTypeLPDDR3,
		memoryTypeLPDDR4,
		memoryTypeLogicalNonVolatileDevice,
		memoryTypeHBM,
		memoryTypeHBM2,
		memoryTypeDDR5,
		memoryTypeLPDDR5,
	}[m]
}

// MemoryType is the type of memory used in this device; see 7.18.2
// for definitions.
func (s MemoryDeviceStructure) MemoryType() MemoryType {
	return MemoryType(s.Formatted[14])
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

const (
	typeDetailAttributeLRDIMM       = "LRDIMM"                    // Bit 15
	typeDetailAttributeUnbuffered   = "Unbuffered (Unregistered)" // Bit 14
	typeDetailAttributeRegistered   = "Registered (Buffered)"     // Bit 13
	typeDetailAttributeNonVolatile  = "Non-volatile"              // Bit 12
	typeDetailAttributeCacheDRAM    = "Cache DRAM"                // Bit 11
	typeDetailAttributeWindowDRAM   = "Window DRAM"               // Bit 10
	typeDetailAttributeEDO          = "EDO"                       // Bit 9
	tTypeDetailAttributeCMOS        = "CMOS"                      // Bit 8
	typeDetailAttributeSynchronous  = "Synchronous"               // Bit 7
	typeDetailAttributeRAMBUS       = "RAMBUS"                    // Bit 6
	typeDetailAttributePseudoStatic = "Pseudo-static"             // Bit 5
	typeDetailAttributeStaticColumn = "Static column"             // Bit 4
	typeDetailAttributeFastPaged    = "Fast-paged"                // Bit 3
	typeDetailAttributeUnknown      = "Unknown"                   // Bit 2
	typeDetailAttributeOther        = "Other"                     // Bit 1
	typeDetailAttributeReserved     = "Reserved"                  // Bit 0
)

func (t TypeDetailAttribute) String() string {
	return [...]string{
		typeDetailAttributeLRDIMM,
		typeDetailAttributeUnbuffered,
		typeDetailAttributeRegistered,
		typeDetailAttributeNonVolatile,
		typeDetailAttributeCacheDRAM,
		typeDetailAttributeWindowDRAM,
		typeDetailAttributeEDO,
		tTypeDetailAttributeCMOS,
		typeDetailAttributeSynchronous,
		typeDetailAttributeRAMBUS,
		typeDetailAttributePseudoStatic,
		typeDetailAttributeStaticColumn,
		typeDetailAttributeFastPaged,
		typeDetailAttributeUnknown,
		typeDetailAttributeOther,
		typeDetailAttributeReserved,
	}[t]
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

// TypeDetail is additional detail on the memory device type; see
// 7.18.3 for definitions.
func (s MemoryDeviceStructure) TypeDetail() TypeDetail {
	return TypeDetail(binary.LittleEndian.Uint16(s.Formatted[15:17]))
}

// MemoryDeviceSpeed represents the SMBIOS memory device speed.
type MemoryDeviceSpeed uint16

func (m MemoryDeviceSpeed) String() string {
	return fmt.Sprintf("%d MT/s", m)
}

// Speed identifies the maximum capable speed of the
// device, in megatransfers per second (MT/s). See
// 7.18.4 for details.
// 0000h = the speed is unknown
// FFFFh = the speed is 65,535 MT/s or greater,
// and the actual speed is stored in the Extended
// Speed field.
func (s MemoryDeviceStructure) Speed() MemoryDeviceSpeed {
	return MemoryDeviceSpeed(binary.LittleEndian.Uint16(s.Formatted[17:19]))
}

// Manufacturer is the string number for the manufacturer of this memory
// device.
func (s MemoryDeviceStructure) Manufacturer() string {
	return get(s.Structure, 2)
}

// SerialNumber is the string number for the serial number of this memory
// device.
// This value is set by the manufacturer and normally
// is not changeable.
func (s MemoryDeviceStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

// AssetTag is the string number for the asset tag of this memory
// device.
func (s MemoryDeviceStructure) AssetTag() string {
	return get(s.Structure, 4)
}

// PartNumber is the string number for the part number of this memory
// device.
// This value is set by the manufacturer and normally
// is not changeable.
func (s MemoryDeviceStructure) PartNumber() string {
	return get(s.Structure, 5)
}

// Attributes returns the memory device attributes.
func (s MemoryDeviceStructure) Attributes() {}

// ExtendedSize returns the memory device extended size.
func (s MemoryDeviceStructure) ExtendedSize() {}

// ConfiguredMemorySpeed represents the SMBIOS memory device configured speed.
type ConfiguredMemorySpeed uint16

func (c ConfiguredMemorySpeed) String() string {
	return fmt.Sprintf("%d MT/s", c)
}

// ConfiguredMemorySpeed identifies the configured speed of the memory
// device, in megatransfers per second (MT/s). See
// 7.18.4 for details.
// 0000h = the speed is unknown
// FFFFh = the speed is 65,535 MT/s or greater,
// and the actual speed is stored in the Extended
// Configured Memory Speed field.
func (s MemoryDeviceStructure) ConfiguredMemorySpeed() ConfiguredMemorySpeed {
	return ConfiguredMemorySpeed(binary.LittleEndian.Uint16(s.Formatted[28:30]))
}

// MemoryDeviceVoltage represents the SMBIOS memory device voltage.
type MemoryDeviceVoltage uint16

func (m MemoryDeviceVoltage) String() string {
	if m == 0 {
		return "Unknown"
	}

	return fmt.Sprintf("%.2f V", float32(m)/1000)
}

// MinimumVoltage is the minimum operating voltage for this device, in
// millivolts
// If the value is 0, the voltage is unknown.
func (s MemoryDeviceStructure) MinimumVoltage() MemoryDeviceVoltage {
	return MemoryDeviceVoltage(binary.LittleEndian.Uint16(s.Formatted[30:32]))
}

// MaximumVoltage is the maximum operating voltage for this device, in
// millivolts
// If the value is 0, the voltage is unknown.
func (s MemoryDeviceStructure) MaximumVoltage() MemoryDeviceVoltage {
	return MemoryDeviceVoltage(binary.LittleEndian.Uint16(s.Formatted[32:34]))
}

// ConfiguredVoltage is the configured voltage for this device, in millivolts
// If the value is 0, the voltage is unknown.
func (s MemoryDeviceStructure) ConfiguredVoltage() MemoryDeviceVoltage {
	return MemoryDeviceVoltage(binary.LittleEndian.Uint16(s.Formatted[34:36]))
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
