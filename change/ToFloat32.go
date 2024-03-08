package change

import (
	"math"
)

func ToFloat32[FromType Simple | SimplePointor](data FromType) (float32, error) {
	f64, err := ToFloat64(data)
	if err != nil {
		return 0, err
	}
	if f64 > math.MaxFloat32 || f64*-1 > math.SmallestNonzeroFloat32 {
		return 0, ErrRange
	}
	return float32(f64), err
}
