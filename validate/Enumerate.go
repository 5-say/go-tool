package validate

// Enumerate .. 校验枚举值
func Enumerate[T string |

	int | int8 | int16 | int32 | int64 |

	uint | uint8 | uint16 | uint32 | uint64 |

	float32 | float64,

](data T, enumerate []T) bool {
	for _, v := range enumerate {
		if v == data {
			return true
		}
	}
	return false
}
