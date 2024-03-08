package change

func To[ToType Simple, FromType Simple | SimplePointor](data FromType) (result ToType, err error) {
	switch any(result).(type) {
	case uint32:
		result, err := ToUint32(data)
		return any(result).(ToType), err
	}
	return result, nil
}

func ForceTo[ToType Simple, FromType Simple | SimplePointor](data FromType, defaultValue ToType) ToType {
	result, err := To[ToType](data)
	if err != nil {
		return defaultValue
	}
	return result
}
