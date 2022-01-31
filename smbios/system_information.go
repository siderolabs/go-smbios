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

// SystemInformation represents the SMBIOS system information.
//
//nolint:govet
type SystemInformation struct {
	// Manufacturer returns the system manufacturer.
	Manufacturer string
	// ProductName returns the system product name.
	ProductName string
	// Version returns the system version.
	Version string
	// SerialNumber returns the system serial number.
	SerialNumber string
	// UUID returns the system UUID.
	UUID string
	// WakeUpType identifies the event that caused the system to
	// power up. See 7.2.2.
	WakeUpType WakeUpType
	// SKUNumber returns the system SKU number.
	SKUNumber string
	// Family returns the system family.
	Family string
}

// NewSystemInformation initializes and returns a new `SystemInformation`.
func NewSystemInformation(s *smbios.Structure, v Version) *SystemInformation {
	uuidString := ""

	uid, err := GetUUID(v, s)
	if err == nil {
		str := uid.String()
		uuidString = str
	}

	return &SystemInformation{
		Manufacturer: GetStringOrEmpty(s, 0x04),
		ProductName:  GetStringOrEmpty(s, 0x05),
		Version:      GetStringOrEmpty(s, 0x06),
		SerialNumber: GetStringOrEmpty(s, 0x07),
		UUID:         uuidString,
		WakeUpType:   WakeUpType(GetByte(s, 0x18)),
		SKUNumber:    GetStringOrEmpty(s, 0x19),
		Family:       GetStringOrEmpty(s, 0x1A),
	}
}

// WakeUpType defines the Wake-up type enum.
type WakeUpType int

const (
	// WakeUpTypeReserved is a wakeup type.
	WakeUpTypeReserved WakeUpType = iota
	// WakeUpTypeOther is a wakeup type.
	WakeUpTypeOther
	// WakeUpTypeUnknown is a wakeup type.
	WakeUpTypeUnknown
	// WakeUpTypeAPMTimer is a wakeup type.
	WakeUpTypeAPMTimer
	// WakeUpTypeModemRing is a wakeup type.
	WakeUpTypeModemRing
	// WakeUpTypeLANRemote is a wakeup type.
	WakeUpTypeLANRemote
	// WakeUpTypePowerSwitch is a wakeup type.
	WakeUpTypePowerSwitch
	// WakeUpTypePCIPME is a wakeup type.
	WakeUpTypePCIPME
	// WakeUpTypeACPowerRestored is a wakeup type.
	WakeUpTypeACPowerRestored
)

func (w WakeUpType) String() string {
	switch w {
	case WakeUpTypeReserved:
		return _Reserved
	case WakeUpTypeOther:
		return _Other
	case WakeUpTypeUnknown:
		return _Unknown
	case WakeUpTypeAPMTimer:
		return "APM Timer"
	case WakeUpTypeModemRing:
		return "Modem Ring"
	case WakeUpTypeLANRemote:
		return "LAN Remote"
	case WakeUpTypePowerSwitch:
		return "Power Switch"
	case WakeUpTypePCIPME:
		return "PCI PME#"
	case WakeUpTypeACPowerRestored:
		return "AC Power Restored"
	}

	return _Unknown
}

// GetUUID returns the system Universal Unique ID number.
// Return middle endian only if SMBIOS version >= 2.6.
// Reference: http://dnaeon.github.io/convert-big-endian-uuid-to-middle-endian/
func GetUUID(v Version, s *smbios.Structure) (uid uuid.UUID, err error) {
	var b []byte
	if v.Major >= 3 || (v.Major == 2 && v.Minor >= 6) {
		b, err = toMiddleEndian(s.Formatted)
		if err != nil {
			return uid, fmt.Errorf("failed to convert to middle endian: %w", err)
		}
	} else {
		b, err = toBigEndian(s.Formatted)
		if err != nil {
			return uid, fmt.Errorf("failed to convert to big endian: %w", err)
		}
	}

	uid, err = uuid.FromBytes(b)
	if err != nil {
		return uid, fmt.Errorf("invalid GetUUID: %w", err)
	}

	return uid, nil
}

func toBigEndian(formatted []byte) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, 16))

	if err := binary.Write(buf, binary.BigEndian, formatted[4:20]); err != nil {
		return nil, err
	}

	b = buf.Bytes()

	return b, nil
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
