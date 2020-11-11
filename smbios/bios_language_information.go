// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// BIOSLanguageInformationStructure represents the SMBIOS BIOS language information structure.
type BIOSLanguageInformationStructure struct {
	smbios.Structure
}

// BIOSLanguageInformation returns a `BIOSLanguageInformationStructure`.
func (s SMBIOS) BIOSLanguageInformation() BIOSLanguageInformationStructure {
	return s.BIOSLanguageInformationStructure
}

// CurrentLanguage returns the current language.
func (s BIOSLanguageInformationStructure) CurrentLanguage() string {
	return get(s.Structure, 0)
}
