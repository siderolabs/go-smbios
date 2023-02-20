// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/siderolabs/go-smbios/smbios/internal/github.com/digitalocean/go-smbios/smbios"

// ProcessorInformation represents the SMBIOS process information.
//
//nolint:govet
type ProcessorInformation struct {
	// SocketDesignation returns the processor socket designation.
	SocketDesignation string
	// ProcessorManufacturer returns the processor manufacturer.
	ProcessorManufacturer string
	// ProcessorVersion returns the processor version.
	ProcessorVersion string
	// MaxSpeed returns the maximum speed of the processor,
	// in MHz, supported by the system for this processor socket.
	// If the value is unknown, the field is set to 0.
	// NOTE: This field identifies a capability for the system,
	// not the processor itself.
	MaxSpeed uint16
	// CurrentSpeed returns the current speed of the processor.
	// in MHz, supported by the system for this processor socket.
	// NOTE: This field identifies the processor's speed at
	// system boot, and the Processor ID field implies the
	// processor's additional speed characteristics (that is,
	// single speed or multiple speed).
	CurrentSpeed uint16
	// ProcessorStatus returns the processor status.
	Status ProcessorStatus
	// SerialNumber returns the processor serial number.
	SerialNumber string
	// AssetTag returns the processor asset tag.
	AssetTag string
	// PartNumber returns the processor part number.
	PartNumber string
	// CoreCount returns the processor's number of cores.
	CoreCount uint8
	// CoreEnabled returns the processor's number of enabled cores.
	CoreEnabled uint8
	// ThreadCount returns the processor's number of threads.
	ThreadCount uint8
}

// NewProcessorInformation initializes and returns a new `ProcessorInformation`.
func NewProcessorInformation(s *smbios.Structure) *ProcessorInformation {
	return &ProcessorInformation{
		SocketDesignation:     GetStringOrEmpty(s, 0x04),
		ProcessorManufacturer: GetStringOrEmpty(s, 0x07),
		ProcessorVersion:      GetStringOrEmpty(s, 0x10),
		MaxSpeed:              GetWord(s, 0x14),
		CurrentSpeed:          GetWord(s, 0x16),
		Status:                ProcessorStatus(GetByte(s, 0x18)),
		SerialNumber:          GetStringOrEmpty(s, 0x20),
		AssetTag:              GetStringOrEmpty(s, 0x21),
		PartNumber:            GetStringOrEmpty(s, 0x22),
		CoreCount:             GetByte(s, 0x23),
		CoreEnabled:           GetByte(s, 0x24),
		ThreadCount:           GetByte(s, 0x25),
	}
}

// ProcessorStatus represents the processor status.
type ProcessorStatus int

// SocketPopulated returns true if the SMBIOS processor socket is populated.
func (s ProcessorStatus) SocketPopulated() bool {
	return IsNthBitSet(int(s), 6)
}
