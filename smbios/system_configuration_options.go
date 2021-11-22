// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import "github.com/digitalocean/go-smbios/smbios"

// SystemConfigurationOptionsStructure  represents the SMBIOS system configuration options structure.
type SystemConfigurationOptionsStructure struct {
	*smbios.Structure
}

// SystemConfigurationOptions returns a `SystemConfigurationOptionsStructure`.
func (s *SMBIOS) SystemConfigurationOptions() SystemConfigurationOptionsStructure {
	return s.SystemConfigurationOptionsStructure
}
