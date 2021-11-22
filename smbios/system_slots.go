// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// SystemSlotsStructure represents the SMBIOS systems slots structure.
type SystemSlotsStructure struct {
	*smbios.Structure
}

// SystemSlots returns a `SystemSlotsStructure`.
func (s *SMBIOS) SystemSlots() SystemSlotsStructure {
	return s.SystemSlotsStructure
}

// SlotDesignation returns the slot designation.
func (s SystemSlotsStructure) SlotDesignation() string {
	return get(s.Structure, 0)
}
