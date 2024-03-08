package change

import (
	"math"
)

func ToInt32[FromType Simple | SimplePointor](data FromType) (int32, error) {
	i64, err := ToInt64(data)
	if err != nil {
		return 0, err
	}
	if i64 > math.MaxInt32 || i64 < math.MinInt32 {
		return 0, ErrRange
	}
	return int32(i64), nil
}
