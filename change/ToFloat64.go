package change

import "strconv"

func ToFloat64[FromType Simple | SimplePointor](data FromType) (float64, error) {

	switch v := any(data).(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case *string:
		return strconv.ParseFloat(*v, 64)
	case *float32:
		return float64(*v), nil
	case *float64:
		return *v, nil
	case *int:
		return float64(*v), nil
	case *int8:
		return float64(*v), nil
	case *int16:
		return float64(*v), nil
	case *int32:
		return float64(*v), nil
	case *int64:
		return float64(*v), nil
	case *uint:
		return float64(*v), nil
	case *uint8:
		return float64(*v), nil
	case *uint16:
		return float64(*v), nil
	case *uint32:
		return float64(*v), nil
	case *uint64:
		return float64(*v), nil
	case *bool:
		if *v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, nil
	}
}
