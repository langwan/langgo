package helper_graphic

import (
	"errors"
	"math"
)

// RangeMapper the value in range src maps to the value in range dst
// the range direction of src and dst is allowed to be inconsistent
// value - src value
// ss - src start
// se - src end
// ds - dst start
// de - dst end
func RangeMapper(value, ss, se, ds, de float64) (float64, error) {
	if ss == se || ds == de {
		return 0, errors.New("start is equal end")
	}
	if ss > se && (value > ss || value < se) {
		return 0, errors.New("the value is out of range")
	} else if ss < se && (value > se || value < ss) {
		return 0, errors.New("the value is out of range")
	}
	sr := math.Abs(ss - se)
	dr := math.Abs(ds - de)
	sd := math.Abs(ss - value)
	dd := dr / sr * sd
	if ds > de {
		return ds - dd, nil
	} else {
		return ds + dd, nil
	}
}
