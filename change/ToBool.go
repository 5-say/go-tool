package change

func ToBool[FromType Simple | SimplePointor](data FromType) (to bool, err error) {

	switch v := any(data).(type) {
	case string:
		return v == "true", nil
	case bool:
		return v, nil
	case *string:
		if v == nil {
			return false, ErrNilPointer
		}
		return *v == "true", nil
	case *bool:
		if v == nil {
			return false, ErrNilPointer
		}
		return *v, nil
	default:
		f64, err := ToFloat64(data)
		return f64 != 0, err
	}
}
