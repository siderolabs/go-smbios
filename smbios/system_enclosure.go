// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// SystemEnclosure represents the system enclosure.
type SystemEnclosure struct {
	// Manufacturer returns the system enclosure manufacturer.
	Manufacturer string
	// Version returns the system enclosure version.
	Version string
	// SerialNumber returns the system enclosure serial number.
	SerialNumber string
	// AssetTagNumber returns the system enclosure asset tag number.
	AssetTagNumber string
	// SKUNumber returns the system enclosure SKU number.
	SKUNumber string
}

// NewSystemEnclosure initializes and returns a new `SystemEnclosure`.
func NewSystemEnclosure(s *smbios.Structure) *SystemEnclosure {
	containedElementCount := GetByte(s, 0x13)
	containedElementRecordLength := GetByte(s, 0x14)
	n, m := int(containedElementCount), int(containedElementRecordLength)

	return &SystemEnclosure{
		Manufacturer: GetStringOrEmpty(s, 0x04),
		Version:      GetStringOrEmpty(s, 0x06),
		SerialNumber: GetStringOrEmpty(s, 0x07),
		SKUNumber:    GetStringOrEmpty(s, 0x15+n*m),
	}
}
