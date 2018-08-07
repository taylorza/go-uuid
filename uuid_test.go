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
	"fmt"
	"testing"
)

var testBytes = []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

func TestNewUUID(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.Fatal(err)
	}

	sum := 0
	for i := 0; i < len(uuid); i++ {
		sum += int(uuid[i])
	}
	if sum == 0 {
		t.Fatalf("failed to generate a valid UUID")
	}
}

func TestFromBytes(t *testing.T) {
	uuid, err := FromBytes(testBytes)
	if err != nil {
		t.Fatalf("failed to create UUID from bytes")
	}

	if err := testUUID(uuid); err != nil {
		t.Fatalf("failed to create UUID from bytes correctly")
	}
}

func TestParse(t *testing.T) {
	uuid, err := Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Fatal(err)
	}

	if err := testUUID(uuid); err != nil {
		t.Fatalf("failed to parse UUID string")
	}
}

func TestParseWithBrace(t *testing.T) {
	uuid, err := Parse("{6ba7b810-9dad-11d1-80b4-00c04fd430c8}")
	if err != nil {
		t.Fatalf("faile to parse UUID string with braces : %v", err)
	}

	if err := testUUID(uuid); err != nil {
		t.Fatalf("failed to parse UUID string with braces")
	}
}

func TestString(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.Fatal(err)
	}

	uuidstr := uuid.String()

	uuid2, err := Parse(uuidstr)
	if err != nil {
		t.Fatal(err)
	}

	if uuid != uuid2 {
		t.Fatal("UUID String failed")
	}
}

func TestFormatWithBraces(t *testing.T) {
	uuid, err := Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Fatal(err)
	}

	uuidstr := uuid.Format(WithBraces)

	if uuidstr != "{6ba7b810-9dad-11d1-80b4-00c04fd430c8}" {
		t.Fatal("UUID Format WithBraces failed")
	}
}

func TestFormatUpperCase(t *testing.T) {
	uuid, err := Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Fatal(err)
	}

	uuidstr := uuid.Format(UpperCase)

	if uuidstr != "6BA7B810-9DAD-11D1-80B4-00C04FD430C8" {
		t.Fatal("UUID Format UpperCase failed")
	}
}

func TestFormatWithBracesAndUpperCase(t *testing.T) {
	uuid, err := Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Fatal(err)
	}

	uuidstr := uuid.Format(WithBraces | UpperCase)

	if uuidstr != "{6BA7B810-9DAD-11D1-80B4-00C04FD430C8}" {
		t.Fatal("UUID Format WithBraces and UpperCase failed")
	}
}

func testUUID(uuid UUID) error {
	if uuid[0] != testBytes[0] ||
		uuid[1] != testBytes[1] ||
		uuid[2] != testBytes[2] ||
		uuid[3] != testBytes[3] ||
		uuid[4] != testBytes[4] ||
		uuid[5] != testBytes[5] ||
		uuid[6] != testBytes[6] ||
		uuid[7] != testBytes[7] ||
		uuid[8] != testBytes[8] ||
		uuid[9] != testBytes[9] ||
		uuid[10] != testBytes[10] ||
		uuid[11] != testBytes[11] ||
		uuid[12] != testBytes[12] ||
		uuid[13] != testBytes[13] ||
		uuid[14] != testBytes[14] ||
		uuid[15] != testBytes[15] {
		return fmt.Errorf("uuid did not match reference value")
	}
	return nil
}

var (
	benchmarkedUUID    UUID
	benchmarkedUUIDStr string
)

func BenchmarkNewUUID(b *testing.B) {
	var uuid UUID
	for n := 0; n < b.N; n++ {
		uuid, _ = NewUUID()
	}
	benchmarkedUUID = uuid
}

func BenchmarkNewUUID2(b *testing.B) {
	var uuid UUID
	for n := 0; n < b.N; n++ {
		var myid UUID
		myid, _ = NewUUID()
		uuid = myid
	}
	benchmarkedUUID = uuid
}

func BenchmarkParse(b *testing.B) {
	var uuid UUID
	for n := 0; n < b.N; n++ {
		uuid, _ = Parse("{6ba7b810-9dad-11d1-80b4-00c04fd430c8}")
	}
	benchmarkedUUID = uuid
}

func BenchmarkString(b *testing.B) {
	guid, err := NewUUID()
	if err != nil {
		b.Fatalf("failed to create test UUID")
	}

	var uuidStr string
	for n := 0; n < b.N; n++ {
		uuidStr = guid.String()
	}
	benchmarkedUUIDStr = uuidStr
}

func BenchmarkFormatUpperWithBraces(b *testing.B) {
	guid, err := NewUUID()
	if err != nil {
		b.Fatalf("failed to create test UUID")
	}

	var uuidStr string
	for n := 0; n < b.N; n++ {
		uuidStr = guid.Format(UpperCase | WithBraces)
	}
	benchmarkedUUIDStr = uuidStr
}
