// Copyright 2017-2018 DigitalOcean.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package smbios

// A Header is a Structure's header.
type Header struct {
	Type   uint8
	Length uint8
	Handle uint16
}

// A Structure is an SMBIOS structure.
type Structure struct { //nolint:govet
	Header    Header
	Formatted []byte
	Strings   []string
}
