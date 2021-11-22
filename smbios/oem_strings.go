// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// OEMStringsStructure represents the SMBIOS OEM strings structure.
type OEMStringsStructure struct {
	*smbios.Structure
}

// OEMStrings returns the OEM strings.
func (s SMBIOS) OEMStrings() OEMStringsStructure {
	return s.OEMStringsStructure
}
