// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// PortConnectorInformation represents the port connector information.
type PortConnectorInformation struct {
	// InternalReferenceDesignator returns the port connector internal reference designator.
	InternalReferenceDesignator string
	// ExternalReferenceDesignator returns the port connector external reference designator.
	ExternalReferenceDesignator string
}

// NewPortConnectorInformation initializes and returns a new `PortConnectorInformation`.
func NewPortConnectorInformation(s *smbios.Structure) *PortConnectorInformation {
	return &PortConnectorInformation{
		InternalReferenceDesignator: GetStringOrEmpty(s, 0x04),
		ExternalReferenceDesignator: GetStringOrEmpty(s, 0x06),
	}
}
