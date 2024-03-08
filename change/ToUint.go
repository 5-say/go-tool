package change

import "math"

func ToUint[FromType Simple | SimplePointor](data FromType) (uint, error) {
	ui64, err := ToUint64(data)
	if err != nil {
		return 0, err
	}
	if ui64 > math.MaxUint {
		return 0, ErrRange
	}
	return uint(ui64), nil
}
