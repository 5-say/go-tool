package change

func ToBool[FromType Simple | SimplePointor](data FromType) bool {

	switch v := any(data).(type) {
	case string:
		return v == "true"
	case bool:
		return v
	case *string:
		return *v == "true"
	case *bool:
		return *v
	default:
		f64, _ := ToFloat64(data)
		return f64 != 0
	}
}
