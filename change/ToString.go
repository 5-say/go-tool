package change

import (
	"fmt"
)

func ToString[FromType Simple | SimplePointor](data FromType) (string, error) {
	switch v := any(data).(type) {
	case *string:
		if v == nil {
			return "", ErrNilPointer
		}
	case *float32:
		if v == nil {
			return "", ErrNilPointer
		}
	case *float64:
		if v == nil {
			return "", ErrNilPointer
		}
	case *int:
		if v == nil {
			return "", ErrNilPointer
		}
	case *int8:
		if v == nil {
			return "", ErrNilPointer
		}
	case *int16:
		if v == nil {
			return "", ErrNilPointer
		}
	case *int32:
		if v == nil {
			return "", ErrNilPointer
		}
	case *int64:
		if v == nil {
			return "", ErrNilPointer
		}
	case *uint:
		if v == nil {
			return "", ErrNilPointer
		}
	case *uint8:
		if v == nil {
			return "", ErrNilPointer
		}
	case *uint16:
		if v == nil {
			return "", ErrNilPointer
		}
	case *uint32:
		if v == nil {
			return "", ErrNilPointer
		}
	case *uint64:
		if v == nil {
			return "", ErrNilPointer
		}
	case *bool:
		if v == nil {
			return "", ErrNilPointer
		}
	}
	return fmt.Sprint(data), nil
}
