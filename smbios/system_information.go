// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package smbios

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/digitalocean/go-smbios/smbios"
	"github.com/google/uuid"
)

// SystemInformationStructure represents the SMBIOS system information structure.
type SystemInformationStructure struct {
	smbios.Structure
}

// WakeUp defines the Wake-up Type enum.
type WakeUp int

const (
	// TypeReserved is a wakeup type.
	TypeReserved WakeUp = iota
	// TypeOther is a wakeup type.
	TypeOther
	// TypeUnknown is a wakeup type.
	TypeUnknown
	// TypeAPMTimer is a wakeup type.
	TypeAPMTimer
	// TypeModemRing is a wakeup type.
	TypeModemRing
	// TypeLANRemote is a wakeup type.
	TypeLANRemote
	// TypePowerSwitch is a wakeup type.
	TypePowerSwitch
	// TypePCIPME is a wakeup type.
	TypePCIPME
	// TypeACPowerRestored is a wakeup type.
	TypeACPowerRestored
)

const (
	typeReserved        = "Reserved"
	typeOther           = "Other"
	typeUnknown         = "Unknown"
	typeAPMTimer        = "APM Timer"
	typeModemRing       = "Modem Ring"
	typeLANRemote       = "LAN Remote"
	typePowerSwitch     = "Power Switch"
	typePCIPME          = "PCI PME#"
	typeACPowerRestored = "AC Power Restored"
)

func (w WakeUp) String() string {
	return [...]string{
		typeReserved,
		typeOther,
		typeUnknown,
		typeAPMTimer,
		typeModemRing,
		typeLANRemote,
		typePowerSwitch,
		typePCIPME,
		typeACPowerRestored,
	}[w]
}

// SystemInformation returns a `SystemInformationStructure`.
func (s SMBIOS) SystemInformation() SystemInformationStructure {
	return s.SystemInformationStructure
}

// Manufacturer returns the system manufacturer.
func (s SystemInformationStructure) Manufacturer() string {
	return get(s.Structure, 0)
}

// ProductName returns the system product name.
func (s SystemInformationStructure) ProductName() string {
	return get(s.Structure, 1)
}

// Version returns the system version.
func (s SystemInformationStructure) Version() string {
	return get(s.Structure, 2)
}

// SerialNumber returns the system serial number.
func (s SystemInformationStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

// SKUNumber returns the system SKU number.
func (s SystemInformationStructure) SKUNumber() string {
	return get(s.Structure, 4)
}

// Family returns the system family.
func (s SystemInformationStructure) Family() string {
	return get(s.Structure, 5)
}

// WakeUpType identifies the event that caused the system to
// power up. See 7.2.2.
func (s SystemInformationStructure) WakeUpType() string {
	return WakeUp(s.Formatted[20]).String()
}

// UUID returns the system UUID.
func (s SystemInformationStructure) UUID() (uid uuid.UUID, err error) {
	b, err := toMiddleEndian(s.Formatted)
	if err != nil {
		return uid, fmt.Errorf("failed to convernt to middle endian: %w", err)
	}

	uid, err = uuid.FromBytes(b)
	if err != nil {
		return uid, fmt.Errorf("invalid UUID: %w", err)
	}

	// TODO(andrewrynhard): Return middle endian only if SBIOS version > 2.6.
	// Reference: http://dnaeon.github.io/convert-big-endian-uuid-to-middle-endian/

	return uid, nil
}

func toMiddleEndian(formatted []byte) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, 16))

	timeLow := binary.BigEndian.Uint32(formatted[4:8])
	if err := binary.Write(buf, binary.LittleEndian, timeLow); err != nil {
		return nil, err
	}

	timeMid := binary.BigEndian.Uint16(formatted[8:10])
	if err := binary.Write(buf, binary.LittleEndian, timeMid); err != nil {
		return nil, err
	}

	timeHigh := binary.BigEndian.Uint16(formatted[10:12])
	if err := binary.Write(buf, binary.LittleEndian, timeHigh); err != nil {
		return nil, err
	}

	clockSeqHi := formatted[12:13][0]
	if err := binary.Write(buf, binary.BigEndian, clockSeqHi); err != nil {
		return nil, err
	}

	clockSeqLow := formatted[13:14][0]
	if err := binary.Write(buf, binary.BigEndian, clockSeqLow); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, formatted[14:20]); err != nil {
		return nil, err
	}

	b = buf.Bytes()

	return b, nil
}
