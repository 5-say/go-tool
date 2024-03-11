package change

import (
	"errors"
	"strconv"
)

var ErrRange = strconv.ErrRange
var ErrNilPointer = errors.New("value is nil pointer")

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
