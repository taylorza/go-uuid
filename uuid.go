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
	"strings"
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

// New generates and returns a new UUID
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
	parts := strings.Split(value, "-")

	if len(parts) != 5 ||
		len(parts[0]) != 8 ||
		len(parts[1]) != 4 ||
		len(parts[2]) != 4 ||
		len(parts[3]) != 4 ||
		len(parts[4]) != 12 {
		return uuid, fmt.Errorf("invalid UUID string format")
	}

	hexstr := strings.Join(parts, "")
	uuidBytes, err := hex.DecodeString(hexstr)

	if err != nil {
		return uuid, err
	}

	copy(uuid[0:], uuidBytes)
	return uuid, nil
}

// String returns the UUID formated as a string
func (uuid UUID) String() string {
	return uuid.Format(None)
}

// Format returns the UUID formated as a string according to the format flags specied
func (uuid UUID) Format(flags FmtFlags) string {
	const (
		defaultFmtStr   = "%08x-%04x-%04x-%04x-%12x"
		uppercaseFmtStr = "%08X-%04X-%04X-%04X-%12X"
	)

	var fmtstr = defaultFmtStr
	if flags&UpperCase == UpperCase {
		fmtstr = uppercaseFmtStr
	}

	if flags&WithBraces == WithBraces {
		fmtstr = "{" + fmtstr + "}"
	}

	return fmt.Sprintf(fmtstr,
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16])
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
