// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// BaseboardInformationStructure represents the SMBIOS baseboard information structure.
type BaseboardInformationStructure struct {
	*smbios.Structure
}

// BaseboardInformation returns a `BaseboardInformationStructure`.
func (s *SMBIOS) BaseboardInformation() BaseboardInformationStructure {
	return s.BaseboardInformationStructure
}

// Manufacturer returns the baseboard manufacturer.
func (s BaseboardInformationStructure) Manufacturer() string {
	return get(s.Structure, 0)
}

// Product returns the baseboard product.
func (s BaseboardInformationStructure) Product() string {
	return get(s.Structure, 1)
}

// Version returns the baseboard version.
func (s BaseboardInformationStructure) Version() string {
	return get(s.Structure, 2)
}

// SerialNumber returns the baseboard serial number.
func (s BaseboardInformationStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

// AssetTag returns the baseboard asset tag.
func (s BaseboardInformationStructure) AssetTag() string {
	return get(s.Structure, 4)
}

// LocationInChassis returns the number of a null-terminated string that
// describes this board's location within the chassis referenced by the
// Chassis Handle (described below in this table)
// NOTE: This field supports a CIM_Container class mapping where:
// 	- LocationWithinContainer is this field.
// 	- GroupComponent is the chassis referenced by Chassis Handle.
// 	- PartComponent is this baseboard
func (s BaseboardInformationStructure) LocationInChassis() string {
	return get(s.Structure, 5)
}
