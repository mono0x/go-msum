package msum

import "math"

// http://code.activestate.com/recipes/393090/

func Sum(values []float64) float64 {
	var partials []float64
	for _, x := range values {
		i := 0
		for _, y := range partials {
			if math.Abs(x) < math.Abs(y) {
				tmp := x
				x = y
				y = tmp
			}
			hi := x + y
			lo := y - (hi - x)
			if lo != 0.0 {
				partials[i] = lo
				i++
			}
			x = hi
		}
		if i < len(partials) {
			partials[i] = x
			partials = partials[0 : i+1]
		} else {
			partials = append(partials, x)
		}
	}
	var result float64
	for _, x := range partials {
		result += x
	}
	return result
}
