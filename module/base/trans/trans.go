package trans

import "math"

var Pi float64

func init() {
	print("test init\n")
	Pi = 4 * math.Atan(1) // init() function computes Pi
}
