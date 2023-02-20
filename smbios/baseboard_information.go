// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/siderolabs/go-smbios/smbios/internal/github.com/digitalocean/go-smbios/smbios"

// BaseboardInformation represents the SMBIOS baseboard information.
type BaseboardInformation struct {
	// Manufacturer returns the baseboard manufacturer.
	Manufacturer string
	// Product returns the baseboard product.
	Product string
	// Version returns the baseboard version.
	Version string
	// SerialNumber returns the baseboard serial number.
	SerialNumber string
	// AssetTag returns the baseboard asset tag.
	AssetTag string
	// LocationInChassis returns the number of a null-terminated string that
	// describes this board's location within the chassis referenced by the
	// Chassis Handle (described below in this table)
	// NOTE: This field supports a CIM_Container class mapping where:
	// 	- LocationWithinContainer is this field.
	// 	- GroupComponent is the chassis referenced by Chassis Handle.
	// 	- PartComponent is this baseboard
	LocationInChassis string
	// BoardType identifies the type of board. See 7.3.2.
	BoardType BoardType
}

// NewBaseboardInformation initializes and returns a new `BaseboardInformation`.
func NewBaseboardInformation(s *smbios.Structure) *BaseboardInformation {
	return &BaseboardInformation{
		Manufacturer:      GetStringOrEmpty(s, 0x04),
		Product:           GetStringOrEmpty(s, 0x05),
		Version:           GetStringOrEmpty(s, 0x06),
		SerialNumber:      GetStringOrEmpty(s, 0x07),
		AssetTag:          GetStringOrEmpty(s, 0x08),
		LocationInChassis: GetStringOrEmpty(s, 0x0A),
		BoardType:         BoardType(GetByte(s, 0x0D)),
	}
}

// BoardType defines the board type enum.
type BoardType int

const (
	// BoardTypeUnknown is a board type.
	BoardTypeUnknown BoardType = iota
	// BoardTypeOther is a board type.
	BoardTypeOther
	// BoardTypeServerBlade is a board type.
	BoardTypeServerBlade
	// BoardTypeConnectivitySwitch is a board type.
	BoardTypeConnectivitySwitch
	// BoardTypeSystemManagementModule is a board type.
	BoardTypeSystemManagementModule
	// BoardTypeProcessorModule is a board type.
	BoardTypeProcessorModule
	// BoardTypeIOModule is a board type.
	BoardTypeIOModule
	// BoardTypeMemoryModule is a board type.
	BoardTypeMemoryModule
	// BoardTypeDaughterBoard is a board type.
	BoardTypeDaughterBoard
	// BoardTypeMotherboard is a board type.
	BoardTypeMotherboard
	// BoardTypeProcessorMemoryModule is a board type.
	BoardTypeProcessorMemoryModule
	// BoardTypeProcessorIOModule is a board type.
	BoardTypeProcessorIOModule
	// BoardTypeInterconnectBoard is a board type.
	BoardTypeInterconnectBoard
)

func (w BoardType) String() string {
	switch w {
	case BoardTypeUnknown:
		return _Unknown
	case BoardTypeOther:
		return _Other
	case BoardTypeServerBlade:
		return "Server Blade"
	case BoardTypeConnectivitySwitch:
		return "Connectivity Switch"
	case BoardTypeSystemManagementModule:
		return "System Management Module"
	case BoardTypeProcessorModule:
		return "Processor Module"
	case BoardTypeIOModule:
		return "I/O Module"
	case BoardTypeMemoryModule:
		return "Memory Module"
	case BoardTypeDaughterBoard:
		return "Daughter board"
	case BoardTypeMotherboard:
		return "Motherboard (includes processor, memory, and I/O)"
	case BoardTypeProcessorMemoryModule:
		return "Processor/Memory Module"
	case BoardTypeProcessorIOModule:
		return "Processor/IO Module"
	case BoardTypeInterconnectBoard:
		return "Interconnect Board"
	}

	return _Unknown
}
