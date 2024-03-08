package change

import (
	"math"
)

func ToInt16[FromType Simple | SimplePointor](data FromType) (int16, error) {
	i64, err := ToInt64(data)
	if err != nil {
		return 0, err
	}
	if i64 > math.MaxInt16 || i64 < math.MinInt16 {
		return 0, ErrRange
	}
	return int16(i64), nil
}
