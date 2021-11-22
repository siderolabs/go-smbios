// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// PortConnectorInformationStructure represents the port connector information structure.
type PortConnectorInformationStructure struct {
	*smbios.Structure
}

// PortConnectorInformation returns a `PortConnectorInformationStructure`.
func (s *SMBIOS) PortConnectorInformation() PortConnectorInformationStructure {
	return s.PortConnectorInformationStructure
}

// InternalReferenceDesignator returns the port connector internal reference designator.
func (s PortConnectorInformationStructure) InternalReferenceDesignator() string {
	return get(s.Structure, 0)
}

// ExternalReferenceDesignator returns the port connector external reference designator.
func (s PortConnectorInformationStructure) ExternalReferenceDesignator() string {
	return get(s.Structure, 2)
}
