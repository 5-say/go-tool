package change

import (
	"math"
)

func ToInt8[FromType Simple | SimplePointor](data FromType) (int8, error) {
	i64, err := ToInt64(data)
	if err != nil {
		return 0, err
	}
	if i64 > math.MaxInt8 || i64 < math.MinInt8 {
		return 0, ErrRange
	}
	return int8(i64), nil
}
