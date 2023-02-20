// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/siderolabs/go-smbios/smbios/internal/github.com/digitalocean/go-smbios/smbios"

// BIOSLanguageInformation represents the SMBIOS BIOS language information.
type BIOSLanguageInformation struct {
	// CurrentLanguage returns the current language.
	CurrentLanguage string
	// InstallableLanguages returns the installable languages.
	InstallableLanguages []string
}

// NewBIOSLanguageInformation initializes and returns a new `BIOSLanguageInformation`.
func NewBIOSLanguageInformation(s *smbios.Structure) *BIOSLanguageInformation {
	return &BIOSLanguageInformation{
		CurrentLanguage:      GetStringOrEmpty(s, 0x15),
		InstallableLanguages: GetStrings(s),
	}
}
