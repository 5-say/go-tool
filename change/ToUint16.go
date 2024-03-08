package change

import "math"

func ToUint16[FromType Simple | SimplePointor](data FromType) (uint16, error) {
	ui64, err := ToUint64(data)
	if err != nil {
		return 0, err
	}
	if ui64 > math.MaxUint16 {
		return 0, ErrRange
	}
	return uint16(ui64), nil
}
