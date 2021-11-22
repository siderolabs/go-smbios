// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios_test

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/talos-systems/go-smbios/smbios"
)

func TestNodeUUID(t *testing.T) {
	s, err := smbios.New()
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) || errors.Is(err, fs.ErrPermission) {
			t.Skip("SMBIOS information is not available")
		}
	}

	require.NoError(t, err)

	_, err = s.SystemInformation().UUID()
	require.NoError(t, err)
}
