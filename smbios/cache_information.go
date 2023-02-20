// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/siderolabs/go-smbios/smbios/internal/github.com/digitalocean/go-smbios/smbios"

// CacheInformation represents the SMBIOS cache information.
type CacheInformation struct {
	// SocketDesignation returns the cache socket designation.
	SocketDesignation string
}

// NewCacheInformation initializes and returns a new `CacheInformation`.
func NewCacheInformation(s *smbios.Structure) *CacheInformation {
	return &CacheInformation{
		SocketDesignation: GetStringOrEmpty(s, 0x04),
	}
}
