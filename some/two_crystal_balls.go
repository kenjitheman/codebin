package some

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {
	stepSize := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := stepSize
	for i < len(breaks) && !breaks[i] {
		i += stepSize
	}
	for j := i - stepSize; j < i; j++ {
		if breaks[j] {
			return j
		}
	}
	return -1
}
