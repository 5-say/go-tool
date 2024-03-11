package change

// 类型转换，转换失败或传入指针为空时，将返回错误
//
//	data FromType
//
// e.g.
//
//	change.To[uint32](data)
func To[ToType Simple, FromType Simple | SimplePointor](data FromType) (to ToType, err error) {
	switch any(to).(type) {
	case bool:
		to, err := ToBool(data)
		return any(to).(ToType), err
	case float32:
		to, err := ToFloat32(data)
		return any(to).(ToType), err
	case float64:
		to, err := ToFloat64(data)
		return any(to).(ToType), err
	case int:
		to, err := ToInt(data)
		return any(to).(ToType), err
	case int8:
		to, err := ToInt8(data)
		return any(to).(ToType), err
	case int16:
		to, err := ToInt16(data)
		return any(to).(ToType), err
	case int32:
		to, err := ToInt32(data)
		return any(to).(ToType), err
	case int64:
		to, err := ToInt64(data)
		return any(to).(ToType), err
	case string:
		to, err := ToString(data)
		return any(to).(ToType), err
	case uint:
		to, err := ToUint(data)
		return any(to).(ToType), err
	case uint8:
		to, err := ToUint8(data)
		return any(to).(ToType), err
	case uint16:
		to, err := ToUint16(data)
		return any(to).(ToType), err
	case uint32:
		to, err := ToUint32(data)
		return any(to).(ToType), err
	case uint64:
		to, err := ToUint64(data)
		return any(to).(ToType), err
	default:
		return
	}
}
