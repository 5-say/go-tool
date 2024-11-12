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
		return fmt.Sprint(*v), nil
	case *float32:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *float64:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *int:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *int8:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *int16:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *int32:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *int64:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *uint:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *uint8:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *uint16:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *uint32:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *uint64:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	case *bool:
		if v == nil {
			return "", ErrNilPointer
		}
		return fmt.Sprint(*v), nil
	}
	return fmt.Sprint(data), nil
}
