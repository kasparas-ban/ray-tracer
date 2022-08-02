package utils

import (
	"fmt"
	"math"
)

func WriteColor(c Color, samples int) string {
	r := c.X
	g := c.Y
	b := c.Z

	// Divide the color by the number of samples and gamma-correct for gamma=2.0
	scale := 1.0 / float64(samples)
	r = math.Sqrt(r * scale)
	g = math.Sqrt(g * scale)
	b = math.Sqrt(b * scale)

	return fmt.Sprintf("%v %v %v\n",
		int(256*Clamp(r, 0.0, 0.999)),
		int(256*Clamp(g, 0.0, 0.999)),
		int(256*Clamp(b, 0.0, 0.999)),
	)
}
