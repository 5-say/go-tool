package change

import (
	"math"
	"strconv"
)

func ToUint64[FromType Simple | SimplePointor](data FromType) (uint64, error) {

	switch v := any(data).(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case float32:
		if v > math.MaxUint64 || v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case float64:
		if v > math.MaxUint64 || v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case int:
		if v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, ErrRange
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case *string:
		return strconv.ParseUint(*v, 10, 64)
	case *float32:
		if *v > math.MaxUint64 || *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *float64:
		if *v > math.MaxUint64 || *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *int:
		if *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *int8:
		if *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *int16:
		if *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *int32:
		if *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *int64:
		if *v < 0 {
			return 0, ErrRange
		}
		return uint64(*v), nil
	case *uint:
		return uint64(*v), nil
	case *uint8:
		return uint64(*v), nil
	case *uint16:
		return uint64(*v), nil
	case *uint32:
		return uint64(*v), nil
	case *uint64:
		return *v, nil
	case *bool:
		if *v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, nil
	}

}
