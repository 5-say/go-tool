package change

import "errors"

var ErrRange = errors.New("value out of range")

type Simple interface {
	string |
		float32 |
		float64 |
		int | uint |
		int8 | uint8 |
		int16 | uint16 |
		int32 | uint32 |
		int64 | uint64 |
		bool
}

type SimplePointor interface {
	*string |
		*float32 |
		*float64 |
		*int | *uint |
		*int8 | *uint8 |
		*int16 | *uint16 |
		*int32 | *uint32 |
		*int64 | *uint64 |
		*bool
}
