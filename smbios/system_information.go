package smbios

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/digitalocean/go-smbios/smbios"
	"github.com/google/uuid"
)

type SystemInformationStructure struct {
	smbios.Structure
}

func (s Smbios) SystemInformation() SystemInformationStructure {
	return s.SystemInformationStructure
}

func (s SystemInformationStructure) Manufacturer() string {
	return get(s.Structure, 0)
}

func (s SystemInformationStructure) ProductName() string {
	return get(s.Structure, 1)
}

func (s SystemInformationStructure) Version() string {
	return get(s.Structure, 2)
}

func (s SystemInformationStructure) SerialNumber() string {
	return get(s.Structure, 3)
}

func (s SystemInformationStructure) SKUNumber() string {
	return get(s.Structure, 4)
}

func (s SystemInformationStructure) Family() string {
	return get(s.Structure, 5)
}

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

	clockSeqHi := uint8(formatted[12:13][0])
	if err := binary.Write(buf, binary.BigEndian, clockSeqHi); err != nil {
		return nil, err
	}

	clockSeqLow := uint8(formatted[13:14][0])
	if err := binary.Write(buf, binary.BigEndian, clockSeqLow); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, formatted[14:20]); err != nil {
		return nil, err
	}

	b = buf.Bytes()

	return b, nil
}
