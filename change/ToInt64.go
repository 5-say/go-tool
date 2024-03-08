package change

import (
	"math"
	"strconv"
)

func ToInt64[FromType Simple | SimplePointor](data FromType) (int64, error) {

	switch v := any(data).(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case float32:
		if v > math.MaxInt64 || v < math.MinInt64 {
			return 0, ErrRange
		}
		return int64(v), nil
	case float64:
		if v > math.MaxInt64 || v < math.MinInt64 {
			return 0, ErrRange
		}
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case *string:
		return strconv.ParseInt(*v, 10, 64)
	case *float32:
		if *v > math.MaxInt64 || *v < math.MinInt64 {
			return 0, ErrRange
		}
		return int64(*v), nil
	case *float64:
		if *v > math.MaxInt64 || *v < math.MinInt64 {
			return 0, ErrRange
		}
		return int64(*v), nil
	case *int:
		return int64(*v), nil
	case *int8:
		return int64(*v), nil
	case *int16:
		return int64(*v), nil
	case *int32:
		return int64(*v), nil
	case *int64:
		return *v, nil
	case *uint:
		return int64(*v), nil
	case *uint8:
		return int64(*v), nil
	case *uint16:
		return int64(*v), nil
	case *uint32:
		return int64(*v), nil
	case *uint64:
		return int64(*v), nil
	case *bool:
		if *v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, nil
	}
}
