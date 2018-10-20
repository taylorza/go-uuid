# UUID [![Build Status](https://travis-ci.org/taylorza/go-uuid.svg?branch=master)](https://travis-ci.org/taylorza/go-uuid) [![Coverage Status](https://coveralls.io/repos/github/taylorza/go-uuid/badge.svg?branch=master)](https://coveralls.io/github/taylorza/go-uuid?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/taylorza/go-uuid)](https://goreportcard.com/report/github.com/taylorza/go-uuid) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/taylorza/go-uuid) 



Package **uuid** implements version 4 UUID as defined in [RFC 4122](http://tools.ietf.org/html/rfc4122).

The package supports creation of UUID as well as parsing from common string formats to create a corresponding UUID. Currently the package only supports the creation of version 4 UUIDs, which provides uniqueness and a good degree of anonymity.

## Installation

Use the 'go' command:

    $ go get github.com/taylorza/go-uuid

## Examples

```go
    // Create new UUID
    id, err := uuid.NewUUID()
    if err != nil {
        fmt.Println("failed to create uuid : %v", err)
        return
    }

    // Create a UUID from a string
    id, err := Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
    if err != nil {
        fmt.Println("failed to create uuid : %v", err)
        return
    }

    // Create a UUID from an existing byte slice
    // The slice must contain exactly 16 bytes
    uuid, err := FromBytes([]byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8})
```

See [tests](https://github.com/taylorza/go-uuid/blob/master/uuid_test.go) for more examples

## Copyright

Copyright (C)2013-2018 by Chris Taylor (taylorza)
See [LICENSE](https://github.com/taylorza/go-uuid/blob/master/LICENSE)
