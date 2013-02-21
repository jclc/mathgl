package mathgl

import "math"

// Epsilon is some tiny value that determines how precisely equal we want out floats to be
const epsilon float64 = 1e-15

// FloatEqual is a safe utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
// Keep in mind that it expects float32s to be converted to float64s before being passed in, because they have to be converted for Abs anyway
func FloatEqual(a, b float64) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b are 0
		return math.Abs(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return math.Abs(a-b)/(math.Abs(a)+math.Abs(b)) < epsilon
}

// FloatEqualThreshold is a utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
// Keep in mind that it expects float32s to be converted to float64s before being passed in, because they have to be converted for Abs anyway
//
// This differs from FloatEqual in that it lets you pass in your comparison threshold, so that you can adjust the comparison value to your specific needs
func FloatEqualThreshold(a, b, epsilon float64) bool {
	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b is 0
		return math.Abs(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return math.Abs(a-b)/(math.Abs(a)+math.Abs(b)) < epsilon
}

// INTERNAL: Makes sure Scalars of of the right type
func checkType(typ VecType, i interface{}) bool {
	switch typ {
	case INT32:
		_, ok := i.(ScalarInt32)
		return ok
	case UINT32:
		_, ok := i.(ScalarUint32)
		return ok
	case FLOAT32:
		_, ok := i.(ScalarFloat32)
		return ok
	case FLOAT64:
		_, ok := i.(ScalarFloat64)
		return ok
	}

	return false
}
