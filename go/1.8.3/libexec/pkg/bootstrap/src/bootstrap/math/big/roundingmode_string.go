// Do not edit. Bootstrap copy of /private/tmp/go-20170524-61097-1ppddl4/go/src/math/big/roundingmode_string.go

//line /private/tmp/go-20170524-61097-1ppddl4/go/src/math/big/roundingmode_string.go:1
// generated by stringer -type=RoundingMode; DO NOT EDIT

package big

import "fmt"

const _RoundingMode_name = "ToNearestEvenToNearestAwayToZeroAwayFromZeroToNegativeInfToPositiveInf"

var _RoundingMode_index = [...]uint8{0, 13, 26, 32, 44, 57, 70}

func (i RoundingMode) String() string {
	if i+1 >= RoundingMode(len(_RoundingMode_index)) {
		return fmt.Sprintf("RoundingMode(%d)", i)
	}
	return _RoundingMode_name[_RoundingMode_index[i]:_RoundingMode_index[i+1]]
}
