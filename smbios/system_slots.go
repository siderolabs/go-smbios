// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// SystemSlot represents a SMBIOS system slot.
type SystemSlot struct {
	// SlotDesignation returns the slot designation.
	SlotDesignation string
}

// NewSystemSlot initializes and returns a new `SystemSlot`.
func NewSystemSlot(s *smbios.Structure) *SystemSlot {
	return &SystemSlot{
		SlotDesignation: GetStringOrEmpty(s, 0x04),
	}
}
