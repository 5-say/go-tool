package change

// 强制类型转换，转换失败或传入指针为空时，将返回默认值
//
//	data FromType
//
// e.g.
//
//	change.ForceTo[uint32](data, defaultValue)
func ForceTo[ToType Simple, FromType Simple | SimplePointor](data FromType, defaultValue ToType) ToType {
	result, err := To[ToType](data)
	if err != nil {
		return defaultValue
	}
	return result
}

// 强制类型转换，转换失败或传入指针为空时，将返回默认值
//
//	data FromType
//
// e.g.
//
//	change.ForceToPointer[uint32](data, defaultValue)
func ForceToPointer[ToType Simple, FromType Simple | SimplePointor](data FromType, defaultValue ToType) *ToType {
	result, err := To[ToType](data)
	if err != nil {
		return &defaultValue
	}
	return &result
}

// 强制类型转换，转换失败或传入指针为空时，将返回 nil
//
//	data FromType
//
// e.g.
//
//	change.ForceToPointerNil[uint32](data)
func ForceToPointerNil[ToType Simple, FromType Simple | SimplePointor](data FromType) *ToType {
	result, err := To[ToType](data)
	if err != nil {
		return nil
	}
	return &result
}
