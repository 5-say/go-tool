package change

import (
	"math"
)

func ToInt[FromType Simple | SimplePointor](data FromType) (int, error) {
	i64, err := ToInt64(data)
	if err != nil {
		return 0, err
	}
	if i64 > math.MaxInt || i64 < math.MinInt {
		return 0, ErrRange
	}
	return int(i64), nil
}
