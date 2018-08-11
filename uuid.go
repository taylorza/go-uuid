// MIT License
//
// Copyright (c) 2018 Chris Taylor (taylorza)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// UUID is a universally unique identifier
type UUID [16]byte

// FmtFlags represents the flags that can be passed to Format to control the resulting string
type FmtFlags int

const (
	// None implies no special formatting of UUID and revets to using the default
	None FmtFlags = iota
	// WithBraces when passed to Format will format the UUID as string enclosed in braces
	WithBraces = 1 << iota
	// UpperCase when passed to Format will format the UUID using uppercase for the hex characters
	UpperCase
)

// NewUUID generates and returns a new UUID
func NewUUID() (UUID, error) {
	return generateV4()
}

// FromBytes returns a UUID from a byte slice. The slice must be 16 bytes in length
func FromBytes(bytes []byte) (UUID, error) {
	var uuid UUID
	if len(bytes) != 16 {
		return uuid, fmt.Errorf("slice is not 16 bytes")
	}
	copy(uuid[0:], bytes)
	return uuid, nil
}

// Parse parses a formated string and returns the UUID it represets.
// the formated string should conform to the following pattern xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// where x represents a hex digit 0-f
func Parse(value string) (UUID, error) {
	var uuid UUID
	if len(value) != 36 && len(value) != 38 {
		return uuid, fmt.Errorf("string is not the correct length")
	}

	if len(value) == 38 {
		if value[0] != '{' && value[37] != '}' {
			return uuid, fmt.Errorf("invalid UUID string format")
		}
		value = value[1:37]
	}
	if value[8] != '-' ||
		value[13] != '-' ||
		value[18] != '-' ||
		value[23] != '-' {
		return uuid, fmt.Errorf("invalid UUID string format")
	}

	if _, err := hex.Decode(uuid[0:], []byte(value[0:8])); err != nil {
		return uuid, fmt.Errorf("invalid UUID : %v", err)
	}
	if _, err := hex.Decode(uuid[4:], []byte(value[9:13])); err != nil {
		return uuid, fmt.Errorf("invalid UUID : %v", err)
	}
	if _, err := hex.Decode(uuid[6:], []byte(value[14:18])); err != nil {
		return uuid, fmt.Errorf("invalid UUID : %v", err)
	}
	if _, err := hex.Decode(uuid[8:], []byte(value[19:23])); err != nil {
		return uuid, fmt.Errorf("invalid UUID : %v", err)
	}
	if _, err := hex.Decode(uuid[10:], []byte(value[24:36])); err != nil {
		return uuid, fmt.Errorf("invalid UUID : %v", err)
	}

	return uuid, nil
}

// String returns the UUID formated as a string
func (uuid UUID) String() string {
	return uuid.Format(None)
}

// Format returns the UUID formated as a string according to the format flags specied
func (uuid UUID) Format(flags FmtFlags) string {
	var fmtbytes []byte
	var offs = 0
	if flags&WithBraces == WithBraces {
		fmtbytes = make([]byte, 38)
		fmtbytes[0] = '{'
		fmtbytes[37] = '}'
		offs = 1
	} else {
		fmtbytes = make([]byte, 36)
	}

	upper := flags&UpperCase == UpperCase
	tohex(fmtbytes[offs+0:offs+8], uuid[0:4], upper)
	fmtbytes[offs+8] = '-'
	tohex(fmtbytes[offs+9:offs+13], uuid[4:6], upper)
	fmtbytes[offs+13] = '-'
	tohex(fmtbytes[offs+14:offs+18], uuid[6:8], upper)
	fmtbytes[offs+18] = '-'
	tohex(fmtbytes[offs+19:offs+23], uuid[8:10], upper)
	fmtbytes[offs+23] = '-'
	tohex(fmtbytes[offs+24:], uuid[10:], upper)

	return string(fmtbytes)
}

func tohex(dst []byte, src []byte, upper bool) {
	const (
		hexlower = "0123456789abcdef"
		hexupper = "0123456789ABCDEF"
	)
	hexchar := hexlower
	if upper {
		hexchar = hexupper
	}
	for i, b := range src {
		dst[i*2+0] = hexchar[(b&0xf0)>>4]
		dst[i*2+1] = hexchar[b&0x0f]
	}
}

func generateV4() (UUID, error) {
	var uuid UUID

	// Generate random bytes
	if _, err := rand.Read(uuid[0:]); err != nil {
		return uuid, fmt.Errorf("failed to generate random bytes: %v", err)
	}

	// Set the version 4 UUID bits
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // clock_seq_hi_and_reserved - set bits 6,7 to 0,1 respectively
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // time_hi_and_version set bits 12-15 to the version number 4
	return uuid, nil
}
