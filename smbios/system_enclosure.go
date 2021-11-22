// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// SystemEnclosureStructure represents the system enclosure structure.
type SystemEnclosureStructure struct {
	*smbios.Structure
}

// SystemEnclosure returns a `SystemEnclosure`.
func (s SMBIOS) SystemEnclosure() SystemEnclosureStructure {
	return s.SystemEnclosureStructure
}

// Manufacturer returns the system enclosure manufacturer.
func (s SystemEnclosureStructure) Manufacturer() string {
	return get(s.Structure, 0)
}

// Version returns the system enclosure version.
func (s SystemEnclosureStructure) Version() string {
	return get(s.Structure, 1)
}

// SerialNumber returns the system enclosure serial number.
func (s SystemEnclosureStructure) SerialNumber() string {
	return get(s.Structure, 2)
}

// AssetTagNumber returns the system enclosure asset tag number.
func (s SystemEnclosureStructure) AssetTagNumber() string {
	return get(s.Structure, 3)
}

// SKUNumber returns the system enclosure SKU number.
func (s SystemEnclosureStructure) SKUNumber() string {
	return get(s.Structure, 4)
}
