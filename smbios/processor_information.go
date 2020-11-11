// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// ProcessorInformationStructure represents the SMBIOS process information structure.
type ProcessorInformationStructure struct {
	smbios.Structure
}

// ProcessorInformation returns a `ProcessorInformationStructure`.
func (s SMBIOS) ProcessorInformation() ProcessorInformationStructure {
	return s.ProcessorInformationStructure
}

// SocketDesignation returns the processor socket designation.
func (s ProcessorInformationStructure) SocketDesignation() string {
	return get(s.Structure, 0)
}

// ProcessorManufacturer returns the processor manufacturer.
func (s ProcessorInformationStructure) ProcessorManufacturer() string {
	return get(s.Structure, 1)
}

// ProcessorVersion returns the processor version.
func (s ProcessorInformationStructure) ProcessorVersion() string {
	return get(s.Structure, 2)
}

// SerialNumber returns the processor serial number.
func (s ProcessorInformationStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

// AssetTag returns the processor asset tag.
func (s ProcessorInformationStructure) AssetTag() string {
	return get(s.Structure, 4)
}

// PartNumber returns the processor part number.
func (s ProcessorInformationStructure) PartNumber() string {
	return get(s.Structure, 5)
}
