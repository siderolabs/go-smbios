// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/siderolabs/go-smbios/smbios/internal/github.com/digitalocean/go-smbios/smbios"

// BIOSInformation represents the BIOS information.
type BIOSInformation struct {
	// Vendor returns the BIOS vendor.
	Vendor string
	// Version returns the BIOS version.
	Version string
	// ReleaseDate returns the BIOS release date.
	ReleaseDate string
}

// NewBIOSInformation initializes and returns a new `BIOSInformation`.
func NewBIOSInformation(s *smbios.Structure) *BIOSInformation {
	return &BIOSInformation{
		GetStringOrEmpty(s, 0x04),
		GetStringOrEmpty(s, 0x05),
		GetStringOrEmpty(s, 0x08),
	}
}
