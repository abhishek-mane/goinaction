package methods_interfaces

import "errors"

// Reference types in Go are the set of slice, map, channel, interface, and function types
// When you declare a variable from one of these types, the value that’s created is called a header values.
// Technically, a string is also a reference type value. All the different header values
// from the different reference types contain a pointer to an underlying data structure.
// Each reference type also contains a set of unique fields that are used to manage the underlying data structure.
// You never share reference type values because the header value is designed to be copied.
// The header value contains a pointer; therefore, you can pass a copy of any reference type value and share the underlying data structure intrinsically.

// Code snippet shows a type called IP which is declared as a slice of bytes.
// Declaring a type like this is useful when you want to declare behavior around a built-in or reference type

type IP []byte

const IPv4len = 4
const IPv6len = 8

func (ip IP) MarshalText() ([]byte, error) {
	if len(ip) == 0 {
		return []byte(""), nil
	}
	if len(ip) != IPv4len && len(ip) != IPv6len {
		return nil, errors.New("invalid IP address")
	}
	return []byte(ip.String()), nil
}
