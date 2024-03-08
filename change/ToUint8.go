package change

import "math"

func ToUint8[FromType Simple | SimplePointor](data FromType) (uint8, error) {
	ui64, err := ToUint64(data)
	if err != nil {
		return 0, err
	}
	if ui64 > math.MaxUint8 {
		return 0, ErrRange
	}
	return uint8(ui64), nil
}
