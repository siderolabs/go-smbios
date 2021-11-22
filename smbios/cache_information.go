// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// CacheInformationStructure represents the SMBIOS cache information structure.
type CacheInformationStructure struct {
	*smbios.Structure
}

// CacheInformation returns a `CacheInformationStructure`.
func (s *SMBIOS) CacheInformation() CacheInformationStructure {
	return s.CacheInformationStructure
}

// SocketDesignation returns the cache socket designation.
func (s CacheInformationStructure) SocketDesignation() string {
	return get(s.Structure, 0)
}
