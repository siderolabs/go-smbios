// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// BIOSInformationStructure represents the BIOS information structure.
type BIOSInformationStructure struct {
	*smbios.Structure
}

// BIOSInformation returns a `BIOSInformationStructure`.
func (s SMBIOS) BIOSInformation() BIOSInformationStructure {
	return s.BIOSInformationStructure
}

// Vendor returns the BIOS vendor.
func (s BIOSInformationStructure) Vendor() string {
	return get(s.Structure, 0)
}

// Version returns the BIOS version.
func (s BIOSInformationStructure) Version() string {
	return get(s.Structure, 1)
}

// ReleaseDate returns the BIOS release date.
func (s BIOSInformationStructure) ReleaseDate() string {
	return get(s.Structure, 2)
}
