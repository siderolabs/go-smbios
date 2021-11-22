// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// GroupAssociationsStructure represents the SMBIOS group associations structure.
type GroupAssociationsStructure struct {
	*smbios.Structure
}

// GroupAssociations returns a `GroupAssociationsStructure`.
func (s *SMBIOS) GroupAssociations() GroupAssociationsStructure {
	return s.GroupAssociationsStructure
}

// GroupName returns the group name.
func (s GroupAssociationsStructure) GroupName() string {
	return get(s.Structure, 0)
}
