package change

import "math"

func ToUint32[FromType Simple | SimplePointor](data FromType) (uint32, error) {
	ui64, err := ToUint64(data)
	if err != nil {
		return 0, err
	}
	if ui64 > math.MaxUint32 {
		return 0, ErrRange
	}
	return uint32(ui64), nil
}
